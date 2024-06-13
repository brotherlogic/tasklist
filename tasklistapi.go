package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	dspb "github.com/brotherlogic/dstore/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	ghbpb "github.com/brotherlogic/githubridge/proto"
	pb "github.com/brotherlogic/tasklist/proto"
)

var (
	CONFIG_KEY = "github.com/brotherlogic/tasklist/config"

	lists = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tasklist_lists",
	})
	closed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tasklist_closed",
	})

	activeLists = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tasklist_active_lists",
	})
)

func getActiveLists(config *pb.Config) int32 {
	count := int32(0)
	for _, list := range config.GetLists() {
		for _, entry := range list.GetTasks() {
			if entry.GetState() != pb.Task_TASK_COMPLETE {
				count++
				break
			}
		}
	}
	return count
}

func (s *Server) metrics(config *pb.Config) {
	lists.Set(float64(len(config.GetLists())))
	activeLists.Set(float64(getActiveLists(config)))
}

func (s *Server) readConfig(ctx context.Context) (*pb.Config, error) {
	data, err := s.dclient.Read(ctx, &dspb.ReadRequest{Key: CONFIG_KEY})
	if err != nil {
		if status.Code(err) != codes.InvalidArgument {
			return nil, err
		}
		data = &dspb.ReadResponse{}
	}

	config := &pb.Config{}
	err = proto.Unmarshal(data.GetValue().GetValue(), config)

	if err == nil {
		s.metrics(config)
	}

	return config, err
}

func (s *Server) saveConfig(ctx context.Context, config *pb.Config) error {
	data, _ := proto.Marshal(config)
	_, err := s.dclient.Write(ctx, &dspb.WriteRequest{Key: CONFIG_KEY, Value: &anypb.Any{Value: data}})
	return err
}

func (s *Server) RenameJob(ctx context.Context, req *pb.RenameJobRequest) (*pb.RenameJobResponse, error) {
	data, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	changed := false
	for _, list := range data.GetLists() {
		for _, task := range list.GetTasks() {
			if task.GetJob() == req.OldJob {
				task.Job = req.NewJob
				changed = true
			}
		}
	}

	if changed {
		return &pb.RenameJobResponse{}, s.saveConfig(ctx, data)
	}

	return nil, status.Errorf(codes.NotFound, "Could not find instances of job %v", req.OldJob)
}

func (s *Server) GetTaskLists(ctx context.Context, req *pb.GetTaskListsRequest) (*pb.GetTaskListsResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskListsResponse{Lists: config.GetLists()}, nil
}

func (s *Server) AddTaskList(ctx context.Context, req *pb.AddTaskListRequest) (*pb.AddTaskListResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	found := false
	for _, list := range config.GetLists() {
		if list.GetName() == req.GetAdd().GetName() {
			found = true
			break
		}
	}

	if found {
		return nil, status.Errorf(codes.AlreadyExists, "%v already exists", req.GetAdd().GetName())
	}

	if getActiveLists(config) >= 3 {
		return nil, status.Errorf(codes.FailedPrecondition, "You have %v lists running, the limit is 3", len(config.Lists))
	}

	// No check delete issue if we've added one
	if req.GetAdd().GetJob() != "" && req.GetAdd().GetIssueNumber() > 0 {
		s.ghclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{User: "brotherlogic", Repo: req.GetAdd().GetJob(), Id: int64(req.GetAdd().GetIssueNumber())})
	}

	config.Lists = append(config.Lists, req.GetAdd())

	err = s.processTaskLists(ctx, config)
	if err != nil {
		return nil, err
	}

	if config.GetTrackingIssue() > 0 && getActiveLists(config) >= 3 {
		err := s.DeleteIssue(ctx, config.GetTrackingIssue())
		if err != nil {
			return nil, err
		}
		config.TrackingIssue = 0
		err = s.saveConfig(ctx, config)
		if err != nil {
			return nil, err
		}
	}

	return &pb.AddTaskListResponse{}, s.saveConfig(ctx, config)
}

func (s *Server) ChangeUpdate(ctx context.Context, req *pbgh.ChangeUpdateRequest) (*pbgh.ChangeUpdateResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	for _, list := range config.GetLists() {
		for _, task := range list.GetTasks() {
			if task.Job == req.GetIssue().GetService() && task.IssueNumber == req.GetIssue().GetNumber() {
				task.State = pb.Task_TASK_COMPLETE
				s.processTaskLists(ctx, config)
				return &pbgh.ChangeUpdateResponse{}, s.saveConfig(ctx, config)
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "Issue %v/%v was not found", req.GetIssue().GetService(), req.GetIssue().GetNumber())
}

func (s *Server) ValidateTaskLists(ctx context.Context, req *pb.ValidateTaskListsRequest) (*pb.ValidateTaskListsResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	err = s.validateLists(ctx, config)
	if err != nil {
		return &pb.ValidateTaskListsResponse{}, err
	}

	s.CtxLog(ctx, fmt.Sprintf("Running %v active lists", getActiveLists(config)))
	if getActiveLists(config) < 3 {
		issue, err := s.ImmediateIssue(ctx, "You need to add a list", fmt.Sprintf("You only have %v lists currently", getActiveLists(config)), false, false)
		if err != nil {
			return nil, err
		}
		config.TrackingIssue = issue.GetNumber()
		err = s.saveConfig(ctx, config)
		if err != nil {
			return nil, err
		}
	}

	if config.GetTrackingIssue() > 0 && getActiveLists(config) >= 3 {
		err := s.DeleteIssue(ctx, config.GetTrackingIssue())
		if err != nil {
			return nil, err
		}
		config.TrackingIssue = 0
		err = s.saveConfig(ctx, config)
		if err != nil {
			return nil, err
		}
	}

	// Do a best effort save here since processing may well fail
	s.saveConfig(ctx, config)

	err = s.processTaskLists(ctx, config)
	if err != nil {
		return &pb.ValidateTaskListsResponse{}, err
	}

	return &pb.ValidateTaskListsResponse{}, s.saveConfig(ctx, config)
}

func (s *Server) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []*pb.Task
	for _, list := range config.GetLists() {
		found := false
		for _, tag := range req.GetTags() {
			if tag == list.GetTag() {
				found = true
			}
		}

		if found {
			for _, task := range list.GetTasks() {
				if task.GetState() == pb.Task_TASK_IN_PROGRESS {
					tasks = append(tasks, task)
				}
			}
		}
	}

	return &pb.GetTasksResponse{Tasks: tasks}, nil
}

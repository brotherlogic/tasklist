package main

import (
	"fmt"
	"sort"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/tasklist/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
)

func (s *Server) getIssueNumber(ctx context.Context, title, service string) (int32, error) {
	issues, err := s.ghclient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return -1, err
	}

	for _, issue := range issues.GetIssues() {
		if issue.GetService() == service && issue.GetTitle() == title {
			return issue.GetNumber(), nil
		}
	}

	return -1, status.Errorf(codes.NotFound, "Could not find %v/%v", title, service)
}

func (s Server) processTaskLists(ctx context.Context, config *pb.Config) error {
	for _, list := range config.GetLists() {
		sort.SliceStable(list.Tasks, func(i, j int) bool {
			return list.Tasks[i].GetIndex() < list.Tasks[j].GetIndex()
		})

		for _, item := range list.GetTasks() {
			if item.GetState() == pb.Task_UNKNOWN || item.GetState() == pb.Task_TASK_WAITING {
				issue, err := s.ghclient.AddIssue(ctx, &pbgh.Issue{
					Title:       item.GetTitle(),
					Service:     item.GetJob(),
					Body:        item.GetTitle(),
					Subscribers: []string{"tasklist"},
				})
				if err != nil {
					if status.Code(err) == codes.AlreadyExists {
						num, err2 := s.getIssueNumber(ctx, item.GetTitle(), item.GetJob())
						if err2 != nil {
							return fmt.Errorf("unable to get issue number %w from add err %w", err2, err)
						}
						item.State = pb.Task_TASK_IN_PROGRESS
						item.IssueNumber = num
					} else {
						if status.Code(err) != codes.ResourceExhausted {
							s.RaiseIssue("Bad issue add", fmt.Sprintf("Unable to add issue: %v", err))
							return err
						}
					}
				} else {
					item.State = pb.Task_TASK_IN_PROGRESS
					item.IssueNumber = issue.GetNumber()
				}
			}

			// Stop at the first task in progress
			if item.GetState() == pb.Task_TASK_IN_PROGRESS {
				break
			}
		}
	}

	return nil
}

func (s *Server) validateLists(ctx context.Context, config *pb.Config) error {
	for _, list := range config.GetLists() {
		for _, task := range list.GetTasks() {
			if task.State == pb.Task_TASK_IN_PROGRESS {
				found := false
				issues, err := s.ghclient.GetIssues(ctx, &pbgh.GetAllRequest{})
				if err != nil {
					return err
				}
				for _, issue := range issues.GetIssues() {
					if issue.GetService() == task.GetJob() && issue.GetNumber() == task.GetIssueNumber() && issue.State != pbgh.Issue_CLOSED {
						found = true
					}
				}

				if !found {
					task.State = pb.Task_TASK_COMPLETE
				}
			}
		}
	}

	return nil
}

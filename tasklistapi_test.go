package main

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dstore_client "github.com/brotherlogic/dstore/client"
	githubridgeclient "github.com/brotherlogic/githubridge/client"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/tasklist/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.dclient = &dstore_client.DStoreClient{Test: true}
	s.ghclient = githubridgeclient.GetTestClient()
	s.SkipIssue = true
	return s
}

func TestRenameJob(t *testing.T) {
	s := InitTestServer()
	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
		Tasks: []*pb.Task{
			{Title: "test",
				Job: "donkey"},
		},
	}})

	if err != nil {
		t.Fatalf("Error in adding list: %v", err)
	}

	_, err = s.RenameJob(context.Background(), &pb.RenameJobRequest{OldJob: "donkey", NewJob: "zebra"})
	if err != nil {
		t.Fatalf("Bad job rename: %v", err)
	}

	lists, err := s.GetTaskLists(context.Background(), &pb.GetTaskListsRequest{})
	if err != nil {
		t.Fatalf("Bad get: %v", err)
	}

	for _, list := range lists.GetLists() {
		for _, task := range list.GetTasks() {
			if task.GetJob() == "donkey" {
				t.Errorf("Job has not been renamed: %v", task)
			}
		}
	}
}

func TestAddList(t *testing.T) {
	s := InitTestServer()
	s.dclient.ErrorCode = make(map[string]codes.Code)
	s.dclient.ErrorCode[CONFIG_KEY] = codes.InvalidArgument

	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
		Tasks: []*pb.Task{
			{Title: "test"},
		},
	}})
	delete(s.dclient.ErrorCode, CONFIG_KEY)

	if err != nil {
		t.Errorf("Failed to add list: %v", err)
	}

	lists, err := s.GetTaskLists(context.Background(), &pb.GetTaskListsRequest{})
	if err != nil {
		t.Fatalf("COuld not get lists: %v", err)
	}

	if len(lists.GetLists()) == 0 || lists.GetLists()[0].Name != "Test" {
		t.Errorf("Bad list pull: %v", lists)
	}
}

func TestAddListNameClash(t *testing.T) {
	s := InitTestServer()

	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
	}})

	if err != nil {
		t.Errorf("Failed to add list: %v", err)
	}

	_, err = s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
	}})

	if status.Code(err) != codes.AlreadyExists {
		t.Errorf("should have reported already exists: %v", err)
	}
}

func TestAddListFailReads(t *testing.T) {
	s := InitTestServer()
	s.dclient.ErrorCode = make(map[string]codes.Code)
	s.dclient.ErrorCode[CONFIG_KEY] = codes.DataLoss

	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
	}})

	if err == nil {
		t.Errorf("Should have failed: %v", err)
	}

	_, err = s.GetTaskLists(context.Background(), &pb.GetTaskListsRequest{})
	if err == nil {
		t.Errorf("Should have failed on gettasklists: %v", err)
	}
}

func TestMoveToNextItemOnChange(t *testing.T) {
	s := InitTestServer()

	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
		Tasks: []*pb.Task{
			&pb.Task{Index: 1, Title: "Task 1", Job: "home"},
			&pb.Task{Index: 2, Title: "Task 2", Job: "home"},
		},
	}})

	if err != nil {
		t.Fatalf("Bad add: %v", err)
	}

	// Task 1 should have been assigned a number
	lists, err := s.GetTaskLists(context.Background(), &pb.GetTaskListsRequest{})
	if err != nil {
		t.Fatalf("Badd list: %v", err)
	}

	number := int32(0)
	for _, list := range lists.GetLists() {
		for _, task := range list.GetTasks() {
			if task.GetTitle() == "Task 1" {
				number = (task.GetIssueNumber())
			}
		}
	}

	if number == 0 {
		t.Fatalf("Issue was not given a number")
	}

	nnumber := int32(0)
	for _, list := range lists.GetLists() {
		for _, task := range list.GetTasks() {
			if task.GetTitle() == "Task 2" {
				nnumber = (task.GetIssueNumber())
			}
		}
	}

	if nnumber != 0 {
		t.Fatalf("Second task was given a number")
	}

	s.ChangeUpdate(context.Background(), &pbgh.ChangeUpdateRequest{Issue: &pbgh.Issue{Title: "Task 1", Service: "home", Number: number}})

	// Task 2 should have been assigned a number
	lists, err = s.GetTaskLists(context.Background(), &pb.GetTaskListsRequest{})
	if err != nil {
		t.Fatalf("Badd list: %v", err)
	}

	number = int32(0)
	for _, list := range lists.GetLists() {
		for _, task := range list.GetTasks() {
			if task.GetTitle() == "Task 2" {
				number = (task.GetIssueNumber())
			}
		}
	}

	if number == 0 {
		t.Errorf("Task was not updated: %v", lists)
	}
}

func TestEmptyChange(t *testing.T) {
	s := InitTestServer()

	_, err := s.ChangeUpdate(context.Background(), &pbgh.ChangeUpdateRequest{})
	if status.Code(err) != codes.NotFound {
		t.Errorf("Empty change should be not found: %v", err)
	}
}

func TestBadLoad(t *testing.T) {
	s := InitTestServer()
	s.dclient.ErrorCode = map[string]codes.Code{CONFIG_KEY: codes.DataLoss}

	_, err := s.ChangeUpdate(context.Background(), &pbgh.ChangeUpdateRequest{})
	if status.Code(err) != codes.DataLoss {
		t.Errorf("Bad load not plubmed throug: %v", err)
	}
}

func TestValidateFailOnRead(t *testing.T) {
	s := InitTestServer()
	s.dclient.ErrorCode = make(map[string]codes.Code)
	s.dclient.ErrorCode[CONFIG_KEY] = codes.DataLoss

	config, err := s.ValidateTaskLists(context.Background(), &pb.ValidateTaskListsRequest{})
	if err == nil {
		t.Errorf("Should have failed here: %v / %v", config, err)
	}
}

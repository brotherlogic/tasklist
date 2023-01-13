package main

import (
	"testing"

	"golang.org/x/net/context"

	dstore_client "github.com/brotherlogic/dstore/client"

	pb "github.com/brotherlogic/tasklist/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.dclient = &dstore_client.DStoreClient{Test: true}
	return s
}

func TestAddList(t *testing.T) {
	s := InitTestServer()

	_, err := s.AddTaskList(context.Background(), &pb.AddTaskListRequest{Add: &pb.TaskList{
		Name: "Test",
	}})

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

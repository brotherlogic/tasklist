package main

import (
	"log"
	"sort"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/tasklist/proto"

	"golang.org/x/net/context"
)

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
				log.Printf("ADDING %v -> %v", item, err)
				if err != nil {
					return err
				}
				item.State = pb.Task_TASK_IN_PROGRESS
				item.IssueNumber = issue.GetNumber()
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
					if issue.GetService() == task.GetJob() && issue.GetNumber() == task.GetIssueNumber() {
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

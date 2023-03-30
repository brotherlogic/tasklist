package tasklist_client

import (
	"context"

	pbgs "github.com/brotherlogic/goserver"
	pb "github.com/brotherlogic/tasklist/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TasklistClient struct {
	Gs        *pbgs.GoServer
	ErrorCode codes.Code
	Test      bool
}

func (c *TasklistClient) GetTasks(ctx context.Context, in *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	if c.Test {
		if c.ErrorCode != codes.OK {
			return nil, status.Errorf(c.ErrorCode, "Forced Error")
		}
		return &pb.GetTasksResponse{}, nil
	}
	conn, err := c.Gs.FDialServer(ctx, "tasklist")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewTaskListServiceClient(conn)
	return client.GetTasks(ctx, in)
}

package tasktransport

import (
	"context"
	taskpb "grpc_test/proto/task"
)

func (s *service) GetTasks(ctx context.Context, request *taskpb.GetTasksRequest) (*taskpb.GetTasksResponse, error) {
	return &taskpb.GetTasksResponse{
		Tasks: nil,
	}, nil
}

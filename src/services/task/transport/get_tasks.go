package tasktransport

import (
	"context"
	taskpb "grpc_test/proto/task"
	taskbusiness "grpc_test/src/services/task/business"
	transformtask "grpc_test/src/transform"
)

func (s *service) GetTasks(ctx context.Context, request *taskpb.GetTasksRequest) (*taskpb.GetTasksResponse, error) {

	biz := taskbusiness.NewBiz(s.repository)

	tasks, _ := biz.GetTasks(ctx)

	return &taskpb.GetTasksResponse{
		Message: "Get tasks successfully",
		Data:    transformtask.ArrayTaskToArrayTaskPB(tasks),
	}, nil
}

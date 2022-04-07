package tasktransport

import (
	"context"
	taskpb "grpc_test/proto/task"
	taskbusiness "grpc_test/src/services/task/business"
	transformtask "grpc_test/src/transform"
)

func (s *service) CreateTask(ctx context.Context, request *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	biz := taskbusiness.NewBiz(s.repository)

	task, _ := biz.CreateTask(ctx, request.GetName())

	return &taskpb.CreateTaskResponse{
		Message: "Create Task Successfully",
		Data:    transformtask.TaskToTaskPB(*task),
	}, nil

}

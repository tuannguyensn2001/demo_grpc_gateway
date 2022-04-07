package tasktransport

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	taskpb "grpc_test/proto/task"
	taskbusiness "grpc_test/src/services/task/business"
	transformtask "grpc_test/src/transform"
)

func (s *service) UpdateTask(ctx context.Context, request *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	biz := taskbusiness.NewBiz(s.repository)

	task, err := biz.Update(ctx, int(request.Id), request.Name)

	if err != nil {
		return &taskpb.UpdateTaskResponse{}, status.Error(codes.Internal, "Cap nhat that bai")
	}

	return &taskpb.UpdateTaskResponse{
		Message: "Cap nhat thanh cong",
		Data:    transformtask.TaskToTaskPB(*task),
	}, nil

}

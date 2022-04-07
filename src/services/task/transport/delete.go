package tasktransport

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	taskpb "grpc_test/proto/task"
	"grpc_test/src/errors"
	taskbusiness "grpc_test/src/services/task/business"
)

func (s *service) DeleteTask(ctx context.Context, request *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	biz := taskbusiness.NewBiz(s.repository)

	id, err := biz.Delete(ctx, int(request.Id))

	if err != nil {
		switch e := err.(type) {
		case *errors.MyError:
			return &taskpb.DeleteTaskResponse{}, status.Error(e.GetGrpcCode(), e.Message)
		default:
			return &taskpb.DeleteTaskResponse{}, status.Error(codes.Internal, err.Error())
		}

	}

	return &taskpb.DeleteTaskResponse{
		Message: "Xoa thanh cong",
		Data:    int64(id),
	}, nil

}

package tasktransport

import taskpb "grpc_test/proto/task"

type service struct {
	taskpb.UnimplementedTaskServiceServer
}

func NewTaskServer() *service {
	return &service{}
}

package tasktransport

import (
	taskpb "grpc_test/proto/task"
	taskrepository "grpc_test/src/services/task/repository"
)

type service struct {
	taskpb.UnimplementedTaskServiceServer
	repository taskrepository.TaskRepository
}

func NewTaskServer(repository taskrepository.TaskRepository) *service {
	return &service{repository: repository}
}

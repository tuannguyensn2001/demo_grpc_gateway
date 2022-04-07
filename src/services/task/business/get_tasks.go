package taskbusiness

import (
	"context"
	"grpc_test/src/models"
	taskrepository "grpc_test/src/services/task/repository"
)

type biz struct {
	repository store
}

type store = taskrepository.TaskRepository

func NewBiz(store store) *biz {
	return &biz{repository: store}
}

func (b *biz) GetTasks(ctx context.Context) ([]models.Task, error) {

	tasks, err := b.repository.GetAllTasks(ctx)

	return tasks, err

}

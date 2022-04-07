package taskbusiness

import (
	"context"
	"grpc_test/src/models"
)

func (biz *biz) CreateTask(ctx context.Context, name string) (*models.Task, error) {
	id, err := biz.repository.Create(ctx, name)

	if err != nil {
		return nil, err
	}

	task, err := biz.repository.FindById(ctx, int(id))

	if err != nil {
		return nil, err
	}

	return task, nil

}

package taskbusiness

import (
	"context"
	"grpc_test/src/models"
)

func (b *biz) Update(ctx context.Context, id int, name string) (*models.Task, error) {
	err := b.repository.Update(ctx, id, name)

	if err != nil {
		return nil, err
	}

	task, err := b.repository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil

}

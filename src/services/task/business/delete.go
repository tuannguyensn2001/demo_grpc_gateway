package taskbusiness

import (
	"context"
	"grpc_test/src/errors"
	"net/http"
)

func (b *biz) Delete(ctx context.Context, id int) (int, error) {

	task, err := b.repository.FindById(ctx, id)

	if err != nil {
		return -1, errors.NewError("Task nay da ton tai", http.StatusConflict)
	}

	if task != nil {

	}

	err = b.repository.Delete(ctx, id)

	return id, err
}

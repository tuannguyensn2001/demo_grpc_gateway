package taskrepository

import (
	"context"
	"grpc_test/src/models"
)

func (store *store) FindById(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task

	query := "SELECT id,name FROM tasks WHERE id = ?"

	err := store.db.Get(&task, query, id)

	if err != nil {
		return nil, err
	}

	return &task, nil

}

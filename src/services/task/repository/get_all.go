package taskrepository

import (
	"context"
	"grpc_test/src/models"
)

func (store *store) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	query := "SELECT id,name FROM tasks"

	var tasks []models.Task

	err := store.db.Select(&tasks, query)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

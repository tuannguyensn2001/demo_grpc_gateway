package taskrepository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"grpc_test/src/models"
)

type store struct {
	db *sqlx.DB
}

type TaskRepository interface {
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	Create(ctx context.Context, name string) (int64, error)
	FindById(ctx context.Context, id int) (*models.Task, error)
	Update(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
}

func NewTaskRepository(db *sqlx.DB) *store {
	return &store{
		db: db,
	}
}

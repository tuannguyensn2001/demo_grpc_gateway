package taskrepository

import "context"

func (store *store) Create(ctx context.Context, name string) (int64, error) {
	query := "INSERT INTO tasks (name) VALUES (?)"

	result, err := store.db.Exec(query, name)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil

}

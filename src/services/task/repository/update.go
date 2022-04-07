package taskrepository

import "context"

func (store *store) Update(ctx context.Context, id int, name string) error {
	_, err := store.db.Exec("UPDATE tasks SET name = ? WHERE id = ? ", name, id)

	return err

}

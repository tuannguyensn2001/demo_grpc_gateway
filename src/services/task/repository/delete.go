package taskrepository

import "context"

func (store *store) Delete(ctx context.Context, id int) error {
	_, err := store.db.Exec("DELETE FROM tasks WHERE id = ?", id)

	return err
}

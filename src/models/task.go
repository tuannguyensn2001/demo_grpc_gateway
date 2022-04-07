package models

type Task struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

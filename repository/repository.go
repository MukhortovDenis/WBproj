package repository

import "database/sql"

type Authorization interface {
}
type TodoList interface {
}
type TodoItems interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItems
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}

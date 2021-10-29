package repository

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

func NewRepository() *Repository {
	return &Repository{}
}

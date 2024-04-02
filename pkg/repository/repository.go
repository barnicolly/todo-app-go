package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-app-go"
)

type TodoItem interface {
	Create(item todo.TodoItem) (int, error)
	GetAll() ([]todo.TodoItem, error)
	GetById(itemId int) (todo.TodoItem, error)
	Delete(itemId int) error
}

type Repository struct {
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoItem:      NewTodoItemPostgres(db),
	}
}

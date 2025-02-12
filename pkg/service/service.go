package service

import (
	"todo-app-go"
	"todo-app-go/pkg/repository"
)

type TodoItem interface {
	Create(item todo.TodoItem) (int, error)
	GetAll() ([]todo.TodoItem, error)
	GetById(itemId int) (todo.TodoItem, error)
	Delete(itemId int) error
	Update(itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoItem: NewTodoItemService(repos.TodoItem),
	}
}

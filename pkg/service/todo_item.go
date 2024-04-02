package service

import (
	"todo-app-go"
	"todo-app-go/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(item todo.TodoItem) (int, error) {
	return s.repo.Create(item)
}

func (s *TodoItemService) GetAll() ([]todo.TodoItem, error) {
	return s.repo.GetAll()
}

func (s *TodoItemService) GetById(itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(itemId)
}

func (s *TodoItemService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}

func (s *TodoItemService) Update(itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(itemId, input)
}
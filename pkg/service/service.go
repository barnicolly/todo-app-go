package service

import "todo-app-go/pkg/repository"

type TodoItem interface {

}

type Service struct {
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

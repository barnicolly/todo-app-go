package repository

import "github.com/jmoiron/sqlx"

type TodoItem interface {
}

type Repository struct {
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

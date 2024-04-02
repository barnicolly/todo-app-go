package repository

type TodoItem interface {

}

type Repository struct {
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}


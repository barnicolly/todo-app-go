package repository

import (
	"fmt"
	"strings"
	"todo-app-go"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll() ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT id, title, description, done FROM %s`, todoItemsTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT id, title, description, done FROM %s where id = $1`, todoItemsTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, todoItemsTable)
	_, err := r.db.Exec(query, itemId)
	return err
}

func (r *TodoItemPostgres) Update(itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s where ti.id = $%d`, todoItemsTable, setQuery, argId)
	args = append(args, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}

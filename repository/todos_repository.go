package repository

import (
	"crud/model"
	"database/sql"
	"fmt"
	"time"
)

type TodoRepository struct {
	connection *sql.DB
}

func NewTodoRepository(connection *sql.DB) *TodoRepository {
	return &TodoRepository{
		connection: connection,
	}
}

func (tr *TodoRepository) GetAll() (todos []model.Todo, err error) {
	query := "SELECT id, name, done, created_at FROM todos;"

	rows, err := tr.connection.Query(query)

	if err != nil {
		fmt.Errorf("TodoRepository - GetAll - Fetch error")
		return nil, err
	}

	var todo model.Todo

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Done, &todo.CreatedAt)

		if err != nil {
			fmt.Errorf("TodoRepository - GetAll - error mapping the todos")
			return todos, err
		}

		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
}

func (tr *TodoRepository) Create(todoName string) (err error) {
	query, err := tr.connection.Prepare("INSERT INTO todos (name, done, created_at) values ($1, $2, $3) RETURNING id")

	if err != nil {
		return err
	}

	var id int

	err = query.QueryRow(todoName, false, time.Now()).Scan(&id)

	if err != nil {
		fmt.Errorf("TodoRepository - Create - create todo failed")
		return err
	}

	fmt.Printf("TodoRepository - Create - succesfully create Todo with id %d \n", id)

	query.Close()

	return err
}

func (tr *TodoRepository) GetById(id int) (todo *model.Todo, err error) {
	query, err := tr.connection.Prepare("SELECT id, name, done, created_at FROM todos WHERE id = $1")

	if err != nil {
		return todo, err
	}

	todo = &model.Todo{}

	err = query.QueryRow(id).Scan(
		&todo.ID,
		&todo.Name,
		&todo.Done,
		&todo.CreatedAt,
	)

	if err != nil {
		fmt.Printf("TodoRepository - GetById - not found todo with id %d \n", id)
		return nil, nil
	}

	query.Close()

	fmt.Printf("TodoRepository - GetById - succesfully get Todo with id %d \n", id)

	return todo, err
}

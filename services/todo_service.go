package services

import (
	"crud/model"
	"crud/repository"
	"fmt"
	"strconv"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (ts *TodoService) GetAll() (todos []model.Todo, err error) {
	fmt.Println("TodoService - GetAll - fetching todos")

	return ts.todoRepository.GetAll()
}

func (ts *TodoService) Create(todoName string) (err error) {
	fmt.Println("TodoService - Create - creating todos")

	err = ts.todoRepository.Create(todoName)

	if err != nil {
		fmt.Println("TodoService - GetById - failed to create todo")
		return err
	}

	return nil
}

func (ts *TodoService) GetById(id string) (todo *model.Todo, err error) {
	fmt.Printf("TodoService - GetById - parsing todo with id %s \n", id)
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("TodoService - GetById - failed to parse id %s \n", id)
		return nil, err
	}

	todo, err = ts.todoRepository.GetById(parsedId)

	if err != nil {
		return nil, err
	}

	return todo, err
}

func (ts *TodoService) BeDone(id string) (err error) {
	fmt.Println("TodoService - BeDone - Marking a to-do as done")

	parsedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("TodoService - GetById - failed to parse id %s \n", id)
		return err
	}

	err = ts.todoRepository.BeDone(parsedId)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TodoService) RemoveTodo(id string) (err error) {
	fmt.Println("TodoService - RemoveTodo - removing todo")

	parsedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("TodoService - GetById - failed to parse id %s \n", id)
		return err
	}

	err = ts.todoRepository.RemoveTodo(parsedId)

	if err != nil {
		return err
	}

	return nil
}

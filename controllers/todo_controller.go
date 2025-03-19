package controllers

import (
	"crud/controllers/dtos"
	"crud/handles"
	"crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (tc *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := tc.todoService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}

	ctx.JSON(http.StatusOK, todos)
}

func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todoDto dtos.CreateTodoDTO
	err := ctx.BindJSON(&todoDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	err = tc.todoService.Create(todoDto.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (tc *TodoController) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	todo, err := tc.todoService.GetById(id)

	err = handles.TodoErrorHandler(ctx, err, todo)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

package routes

import (
	"crud/controllers"
	"crud/repository"
	"crud/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitTodoRoutes(connection *sql.DB, route *gin.Engine) {
	todoRepository := repository.NewTodoRepository(connection)
	todoService := services.NewTodoService(*todoRepository)
	todoController := controllers.NewTodoController(*todoService)

	groupRoute := route.Group("todo")
	groupRoute.GET("", todoController.GetAllTodos)
	groupRoute.POST("", todoController.CreateTodo)
	groupRoute.GET(":id", todoController.GetTodoById)
	groupRoute.PATCH(":id", todoController.BeDone)
	groupRoute.DELETE(":id", todoController.RemoveTodo)
}

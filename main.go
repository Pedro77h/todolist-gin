package main

import (
	"crud/database"
	"crud/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Get em gin")
}

func main() {
	router := gin.Default()
	db, err := database.ConnectDB()

	if err != nil {
		panic((err))
	}

	routes.InitTodoRoutes(db, router)
	router.Run(":8080")
}

package handles

import (
	"crud/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TodoErrorHandler(ctx *gin.Context, err error, todo *model.Todo) error {
	fmt.Print(todo, err)
	if _, isNumError := err.(*strconv.NumError); isNumError {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error, error to parse id",
		})
		return err
	}

	if err == nil && todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return errors.New("todo not found")
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	return err

}

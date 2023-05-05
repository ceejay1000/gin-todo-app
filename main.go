package main

import (
	"net/http"

	handlers "github.com/ceejay1000/todo-app/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	todosApi := router.Group("/api/v1")

	{
		todosApi.GET("/", func(ctx *gin.Context) {
			// ctx.JSON(http.StatusOK, gin.H{
			// 	"message": "Todos API"
			// })
			ctx.String(http.StatusOK, "Todos APi")
		})

		todosApi.GET("/todos", handlers.GetTodos)

		todosApi.GET("/todos/:id", handlers.GetTodo)

		todosApi.POST("todos", handlers.AddTodo)

		todosApi.PUT("/todos/:id", handlers.UpdateTodo)
	}

	todosApi.DELETE("/todos/:id", handlers.DeleteTodo)

	router.Run(":9091")
}

package handler

import (
	"log"
	"net/http"

	todos "github.com/ceejay1000/todo-app/models"
	"github.com/gin-gonic/gin"
	UUID "github.com/google/uuid"
)

func GetTodos(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, todos.TodosDB)
}

func GetTodo(ctx *gin.Context) {

	id := ctx.Param("id")
	todo := todos.GetTodoById(id)

	if todo != nil {
		ctx.JSON(http.StatusOK, todo)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo with ID '" + id + "' not found",
	})
}

func AddTodo(ctx *gin.Context) {

	newTodo := new(todos.Todo)

	if err := ctx.ShouldBindJSON(newTodo); err != nil {
		log.Println("Error parsing JSON")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid fields passed",
		})
	}

	newTodo.Id = UUID.New().String()

	if !todos.TodosList.AddTodo(newTodo) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Todos already exists",
		})

		return
	}

	ctx.JSON(http.StatusCreated, todos.TodosDB)
}

func UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	if !todos.TodoExistsById(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Todo does not exist",
		})
		return
	}

	updatedTodo := new(todos.Todo)

	if err := ctx.ShouldBindJSON(updatedTodo); err != nil {
		log.Println("Error parsing JSON")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid fields passed",
		})
	}

	updatedTodo.Id = id
	todos.UpdateTodo(updatedTodo)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo Updated Successfully",
	})
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	if !todos.TodoExistsById(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Todo does not exist",
		})
		return
	}

	todos.DeleteTodo(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}

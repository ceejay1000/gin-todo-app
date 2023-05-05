package models

import (
	"strings"
	"time"
)

type Todo struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreatedSince time.Time `json:"created-since"`
}

type Todos []*Todo

var TodosList Todos = Todos{}

var TodosDB map[string]*Todos = map[string]*Todos{
	"todos": &TodosList,
}

func (todos *Todos) AddTodo(todo *Todo) bool {

	newTodoStatus := TodoExists(todo)

	if newTodoStatus {
		return false
	}

	TodosList = append(TodosList, todo)
	return true
}

func TodoExists(newTodo *Todo) bool {

	for _, todo := range TodosList {
		if todo.Id == newTodo.Id || strings.EqualFold(todo.Title, newTodo.Title) {
			return true
		}
	}

	return false
}

func TodoExistsById(id string) bool {

	for _, todo := range TodosList {
		if todo.Id == id {
			return true
		}
	}

	return false
}

func GetTodoById(id string) *Todo {

	for _, todo := range TodosList {
		if todo.Id == id {
			return todo
		}
	}

	return nil
}

func GetTodoByTitle(title string) *Todo {

	for _, todo := range TodosList {
		if strings.EqualFold(todo.Title, title) {
			return todo
		}
	}

	return nil
}

func UpdateTodo(td *Todo) {

	for _, todo := range TodosList {
		if todo.Id == td.Id {
			todo.Title = td.Title
			todo.Body = td.Body
		}
	}
}

func DeleteTodo(id string) {

	for index, todo := range TodosList {
		if todo.Id == id {
			TodosList = append(TodosList[0:index], TodosList[index+1:]...)
		}
	}
}

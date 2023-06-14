package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "0", Item: "Limpar a casa", Completed: false},
	{ID: "1", Item: "Ler um livro", Completed: false},
	{ID: "2", Item: "Formatar computador", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	// Extrair os dados do corpo da requisição e transformar na estrutura newTodo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusOK, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo não foi encontrado")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo não foi encontrado"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func togleTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo não foi encontrado"})
		return
	}

	todo.Completed = !todo.Completed
}

func main() {

	// executar uma instancia default da engine do gin-gonic
	router := gin.Default()

	// Definir os endpoints
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", togleTodo)
	router.Run("localhost:9090")
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Todo represents a todo item
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// todos is a slice to store todo items
var todos []Todo

// GetTodosHandler handles GET requests to fetch all todos
func GetTodosHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

// CreateTodoHandler handles POST requests to create a new todo
func CreateTodoHandler(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	todo.ID = len(todos) + 1
	todos = append(todos, *todo)
	return c.JSON(http.StatusCreated, todo)
}

// UpdateTodoHandler handles PUT requests to update an existing todo
func UpdateTodoHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	for i := range todos {
		if todos[i].ID == idInt {
			todo := new(Todo)
			if err := c.Bind(todo); err != nil {
				return err
			}
			todos[i].Title = todo.Title
			todos[i].Status = todo.Status
			return c.JSON(http.StatusOK, todos[i])
		}
	}
	return c.NoContent(http.StatusNotFound)
}

// DeleteTodoHandler handles DELETE requests to delete an existing todo
func DeleteTodoHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	for i := range todos {
		if todos[i].ID == idInt {
			todos = append(todos[:i], todos[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

// RegisterTodoRoutes registers todo routes with the provided Echo instance
func RegisterTodoRoutes(e *echo.Echo) {
	e.GET("/todos", GetTodosHandler)
	e.POST("/todos", CreateTodoHandler)
	e.PUT("/todos/:id", UpdateTodoHandler)
	e.DELETE("/todos/:id", DeleteTodoHandler)
}

package routes

import (
	"github.com/ferdvtn/todolist/apps/interfaces"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRoute(todosSrv interfaces.TodosService) *echo.Echo {
	todosRoute := NewTodosRoute(todosSrv)

	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/api/v1")

	g.POST("/todos", todosRoute.create)
	g.GET("/todos/:ID", todosRoute.get)
	g.GET("/todos", todosRoute.getAll)

	return e
}

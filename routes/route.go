package routes

import (
	"github.com/ferdvtn/todolist/apps/interfaces"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRoute(todolistSrv interfaces.TodolistService) *echo.Echo {
	todolistRoute := NewTodolistRoute(todolistSrv)

	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/api/v1")

	g.GET("/todolists", todolistRoute.getAll)
	g.GET("/todolists/:ID", todolistRoute.get)

	return e
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/ferdvtn/todolist/apps/dto"
	"github.com/ferdvtn/todolist/apps/interfaces"

	"github.com/labstack/echo"
)

type TodolistRoute struct {
	todolistSrv interfaces.TodolistService
}

func NewTodolistRoute(todolistSrv interfaces.TodolistService) *TodolistRoute {
	return &TodolistRoute{
		todolistSrv: todolistSrv,
	}
}

func (r TodolistRoute) get(ctx echo.Context) error {
	argID := ctx.Param("ID")

	ID, err := strconv.Atoi(argID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	result, err := r.todolistSrv.Get(ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, dto.NewDtoResponse("", dto.ToTodolistDtoResponse(result)))
}

func (r TodolistRoute) getAll(ctx echo.Context) error {
	results, err := r.todolistSrv.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, dto.NewDtoResponse("", dto.ToTodolistsDtoResponse(results)))
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/ferdvtn/todolist/apps/dto"
	"github.com/ferdvtn/todolist/apps/interfaces"

	"github.com/labstack/echo"
)

type TodosRoute struct {
	todosSrv interfaces.TodosService
}

func NewTodosRoute(todosSrv interfaces.TodosService) *TodosRoute {
	return &TodosRoute{
		todosSrv: todosSrv,
	}
}

func (r TodosRoute) create(ctx echo.Context) error {

	return nil
}

func (r TodosRoute) get(ctx echo.Context) error {
	argID := ctx.Param("ID")

	ID, err := strconv.Atoi(argID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	result, err := r.todosSrv.Get(ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, dto.NewDtoResponse("", dto.ToTodosDtoResponse(result)))
}

func (r TodosRoute) getAll(ctx echo.Context) error {
	results, err := r.todosSrv.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewDtoResponse(err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, dto.NewDtoResponse("", dto.ToTodosesDtoResponse(results)))
}

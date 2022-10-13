package main

import (
	"github.com/ferdvtn/todolist/apps/repositories"
	"github.com/ferdvtn/todolist/apps/services"
	"github.com/ferdvtn/todolist/routes"
)

func main() {
	todolistRepo := repositories.NewTodolistMysql()
	todolistSrv := services.NewTodolistService(todolistRepo)

	r := routes.NewRoute(todolistSrv)

	r.Start(":1323")
}

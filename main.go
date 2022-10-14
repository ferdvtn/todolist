package main

import (
	"github.com/ferdvtn/todolist/apps/repositories"
	"github.com/ferdvtn/todolist/apps/services"
	"github.com/ferdvtn/todolist/db"
	"github.com/ferdvtn/todolist/routes"
)

func main() {
	db, err := db.GetDB()
	checkError(err)

	todolistRepo := repositories.NewTodolistMysql(db)
	todolistSrv := services.NewTodolistService(todolistRepo)

	r := routes.NewRoute(todolistSrv)

	r.Start(":1323")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

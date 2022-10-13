package repositories

import (
	"fmt"

	"github.com/ferdvtn/todolist/apps/domains"
)

type TodolistMysql struct{}

func NewTodolistMysql() *TodolistMysql {
	return &TodolistMysql{}
}

func (r TodolistMysql) Create(input domains.Todolist, createdBy string) (domains.Todolist, error) {
	return domains.Todolist{}, nil
}

func (r TodolistMysql) Update(input domains.Todolist, updatedBy string) (domains.Todolist, error) {
	return domains.Todolist{}, nil
}

func (r TodolistMysql) Delete(ID int, deletedBy string) error {
	return nil
}

func (r TodolistMysql) Get(ID int) (domains.Todolist, error) {
	return domains.Todolist{
		ID:          ID,
		Title:       fmt.Sprint("task ", ID),
		Description: fmt.Sprint("Desc of task ", ID),
		Priority:    1,
	}, nil
}

func (r TodolistMysql) GetAll() ([]domains.Todolist, error) {
	return []domains.Todolist{
		{
			ID:          1,
			Title:       "task 1",
			Description: "Desc of task 1",
			Priority:    1,
		},
		{
			ID:          2,
			Title:       "task 2",
			Description: "Desc of task 2",
			Priority:    2,
		},
	}, nil
}

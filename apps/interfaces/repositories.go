package interfaces

import "github.com/ferdvtn/todolist/apps/domains"

type TodolistRepo interface {
	Create(input domains.Todolist, createdBy string) (domains.Todolist, error)
	Update(input domains.Todolist, updatedBy string) (domains.Todolist, error)
	Delete(ID int, deletedBy string) error
	Get(ID int) (domains.Todolist, error)
	GetAll() ([]domains.Todolist, error)
}

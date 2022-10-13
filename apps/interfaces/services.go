package interfaces

import "github.com/ferdvtn/todolist/apps/domains"

type TodolistService interface {
	Create(title, description string, priority int, createdBy string) (domains.Todolist, error)
	Update(title, description string, priority int, updatedBy string) (domains.Todolist, error)
	Delete(ID int, deletedBy string) error
	Get(ID int) (domains.Todolist, error)
	GetAll() ([]domains.Todolist, error)
}

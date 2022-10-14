package interfaces

import "github.com/ferdvtn/todolist/apps/domains"

type TodosRepo interface {
	Create(input domains.Todos, createdBy string) (domains.Todos, error)
	Update(input domains.Todos, updatedBy string) (domains.Todos, error)
	Delete(ID int, deletedBy string) error
	Get(ID int) (domains.Todos, error)
	GetAll() ([]domains.Todos, error)
}

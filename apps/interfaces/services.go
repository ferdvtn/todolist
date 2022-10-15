package interfaces

import "github.com/ferdvtn/todolist/apps/domains"

type TodosService interface {
	Create(title, description string, priority int, createdBy string) (domains.Todos, error)
	Update(ID int, title, description string, priority int, updatedBy string) (domains.Todos, error)
	Delete(ID int, deletedBy string) error
	Get(ID int) (domains.Todos, error)
	GetAll() ([]domains.Todos, error)
}

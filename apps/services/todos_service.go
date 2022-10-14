package services

import (
	"github.com/ferdvtn/todolist/apps/domains"
	"github.com/ferdvtn/todolist/apps/interfaces"
)

type TodosService struct {
	todosRepo interfaces.TodosRepo
}

func NewTodosService(todosRepo interfaces.TodosRepo) *TodosService {
	return &TodosService{
		todosRepo: todosRepo,
	}
}

func (srv TodosService) Create(title, description string, priority int, createdBy string) (domains.Todos, error) {
	arg := domains.Todos{
		Title:       title,
		Description: description,
		Priority:    priority,
	}

	err := arg.IsValid()
	if err != nil {
		return domains.Todos{}, err
	}

	result, err := srv.todosRepo.Create(arg, createdBy)
	if err != nil {
		return domains.Todos{}, err
	}

	return result, nil
}

func (srv TodosService) Update(title, description string, priority int, updatedBy string) (domains.Todos, error) {
	return domains.Todos{}, nil
}

func (srv TodosService) Delete(ID int, deletedBy string) error {
	return nil
}

func (srv TodosService) Get(ID int) (domains.Todos, error) {
	result, err := srv.todosRepo.Get(ID)
	if err != nil {
		return domains.Todos{}, err
	}

	return result, nil
}

func (srv TodosService) GetAll() ([]domains.Todos, error) {
	results, err := srv.todosRepo.GetAll()
	if err != nil {
		return []domains.Todos{}, err
	}

	return results, nil
}

package services

import (
	"github.com/ferdvtn/todolist/apps/domains"
	"github.com/ferdvtn/todolist/apps/interfaces"
)

type TodolistService struct {
	interfaces.TodolistRepo
}

func NewTodolistService(todolistRepo interfaces.TodolistRepo) *TodolistService {
	return &TodolistService{
		TodolistRepo: todolistRepo,
	}
}

func (srv TodolistService) Create(title, description string, priority int, createdBy string) (domains.Todolist, error) {
	return domains.Todolist{}, nil
}

func (srv TodolistService) Update(title, description string, priority int, updatedBy string) (domains.Todolist, error) {
	return domains.Todolist{}, nil
}

func (srv TodolistService) Delete(ID int, deletedBy string) error {
	return nil
}

func (srv TodolistService) Get(ID int) (domains.Todolist, error) {
	result, err := srv.TodolistRepo.Get(ID)
	if err != nil {
		return domains.Todolist{}, err
	}

	return result, nil
}

func (srv TodolistService) GetAll() ([]domains.Todolist, error) {
	results, err := srv.TodolistRepo.GetAll()
	if err != nil {
		return []domains.Todolist{}, err
	}

	return results, nil
}

package dto

import "github.com/ferdvtn/todolist/apps/domains"

type TodolistDtoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func ToTodolistDtoResponse(d domains.Todolist) TodolistDtoResponse {
	return TodolistDtoResponse{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Priority:    d.Priority,
	}
}

func ToTodolistsDtoResponse(ds []domains.Todolist) []TodolistDtoResponse {
	var results []TodolistDtoResponse

	if len(ds) == 0 {
		return results
	}

	for _, d := range ds {
		results = append(results, ToTodolistDtoResponse(d))
	}

	return results
}

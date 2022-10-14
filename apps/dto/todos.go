package dto

import "github.com/ferdvtn/todolist/apps/domains"

type TodosDtoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type TodosDtoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func ToTodosDtoResponse(d domains.Todos) TodosDtoResponse {
	return TodosDtoResponse{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Priority:    d.Priority,
	}
}

func ToTodosesDtoResponse(ds []domains.Todos) []TodosDtoResponse {
	var results []TodosDtoResponse

	if len(ds) == 0 {
		return results
	}

	for _, d := range ds {
		results = append(results, ToTodosDtoResponse(d))
	}

	return results
}

package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ferdvtn/todolist/apps/domains"
)

type TodosMysql struct {
	db *sql.DB
}

func NewTodosMysql(db *sql.DB) *TodosMysql {
	return &TodosMysql{
		db: db,
	}
}

func (repo TodosMysql) Create(input domains.Todos, createdBy string) (domains.Todos, error) {
	return domains.Todos{}, nil
}

func (repo TodosMysql) Update(input domains.Todos, updatedBy string) (domains.Todos, error) {
	return domains.Todos{}, nil
}

func (repo TodosMysql) Delete(ID int, deletedBy string) error {
	return nil
}

func (repo TodosMysql) Get(ID int) (domains.Todos, error) {
	return domains.Todos{
		ID:          ID,
		Title:       fmt.Sprint("task ", ID),
		Description: fmt.Sprint("Desc of task ", ID),
		Priority:    1,
	}, nil
}

func (repo TodosMysql) GetAll() ([]domains.Todos, error) {
	return []domains.Todos{
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

package repositories

import (
	"database/sql"

	"github.com/ferdvtn/todolist/apps/domains"
	"github.com/ferdvtn/todolist/helpers"
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
	result, err := repo.db.Exec(
		"INSERT INTO `todos` (`title`, `description`, `priority`) VALUES (?, ?, ?);",
		input.Title,
		helpers.ToNullString(input.Description),
		input.Priority,
	)
	if err != nil {
		return domains.Todos{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domains.Todos{}, err
	}

	input.ID = int(id)

	return input, nil
}

func (repo TodosMysql) Update(input domains.Todos, updatedBy string) (domains.Todos, error) {
	_, err := repo.db.Exec(
		"UPDATE `todos` SET `title` = ?, `description` = ?, `priority` = ? WHERE `id` = ?;",
		input.Title,
		helpers.ToNullString(input.Description),
		input.Priority,
		input.ID,
	)
	if err != nil {
		return domains.Todos{}, err
	}

	return input, nil
}

func (repo TodosMysql) Delete(ID int, deletedBy string) error {
	_, err := repo.db.Exec("DELETE FROM `todos` WHERE `id` = ?;", ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo TodosMysql) Get(ID int) (domains.Todos, error) {
	var result domains.Todos
	err := repo.db.QueryRow("SELECT `id`, `title`, `description`, `priority` FROM `todos` WHERE `id` = ?;", ID).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Priority,
	)
	if err != nil {
		return domains.Todos{}, err
	}

	return result, nil
}

func (repo TodosMysql) GetAll() ([]domains.Todos, error) {
	var results []domains.Todos

	rows, err := repo.db.Query("SELECT `id`, `title`, `description`, `priority` FROM `todos`;")
	if err != nil {
		return []domains.Todos{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var result domains.Todos
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&result.Description,
			&result.Priority,
		)
		if err != nil {
			return []domains.Todos{}, err
		}

		results = append(results, result)
	}

	return results, nil
}

package repository

import (
	"database/sql"
	"todo-service/internal/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func (r *TodoRepository) Create(todo *models.Todo) error {

	query := `
	INSERT INTO todos(user_id,title,completed)
	VALUES($1,$2,$3)
	RETURNING id`

	return r.DB.QueryRow(
		query,
		todo.UserID,
		todo.Title,
		todo.Completed,
	).Scan(&todo.ID)
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {

	rows, err := r.DB.Query(
		`SELECT id,user_id,title,completed FROM todos`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {

		var t models.Todo

		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Title,
			&t.Completed,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func (r *TodoRepository) GetByID(id int) (*models.Todo, error) {

	var todo models.Todo

	query := `
	SELECT id,user_id,title,completed
	FROM todos
	WHERE id=$1`

	err := r.DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Completed,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Update(todo *models.Todo) error {

	query := `
	UPDATE todos
	SET title=$1, completed=$2
	WHERE id=$3`

	_, err := r.DB.Exec(
		query,
		todo.Title,
		todo.Completed,
		todo.ID,
	)

	return err
}

func (r *TodoRepository) Delete(id int) error {

	_, err := r.DB.Exec(
		`DELETE FROM todos WHERE id=$1`,
		id,
	)

	return err
}

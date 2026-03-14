package service

import (
	"todo-service/internal/models"
	"todo-service/internal/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {

	return s.Repo.Create(todo)
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {

	return s.Repo.GetAll()
}

func (s *TodoService) GetTodo(id int) (*models.Todo, error) {

	return s.Repo.GetByID(id)
}

func (s *TodoService) UpdateTodo(todo *models.Todo) error {

	return s.Repo.Update(todo)
}

func (s *TodoService) DeleteTodo(id int) error {

	return s.Repo.Delete(id)
}

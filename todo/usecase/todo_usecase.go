package usecase

import (
	"fmt"

	"github.com/mrizalr/devcode-todolist/models"
	"github.com/mrizalr/devcode-todolist/todo"
)

type todoUsecase struct {
	todoRepository todo.TodoRepository
}

func NewTodoUsecase(todoRepository todo.TodoRepository) todo.TodoUsecase {
	return &todoUsecase{todoRepository}
}

func (u *todoUsecase) GetAll(activityID interface{}) ([]models.GetTodoResponse, error) {
	todos, err := u.todoRepository.Find(activityID)
	todosResponse := []models.GetTodoResponse{}
	if err != nil {
		return todosResponse, err
	}

	for _, todo := range todos {
		t := models.GetTodoResponse{
			ID:         todo.ID,
			ActivityID: todo.ActivityID,
			Title:      todo.Title,
			IsActive:   todo.IsActive,
			Priority:   todo.Priority,
			CreatedAt:  todo.CreatedAt,
			UpdatedAt:  todo.UpdatedAt,
		}

		todosResponse = append(todosResponse, t)
	}

	return todosResponse, nil
}

func (u *todoUsecase) GetByID(todoID int) (models.GetTodoResponse, error) {
	todo, err := u.todoRepository.FindByID(todoID)
	if err != nil {
		return models.GetTodoResponse{}, err
	}

	return models.GetTodoResponse{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Create(todo models.Todo) (models.GetTodoResponse, error) {
	if (todo == models.Todo{} || todo.Title == "") {
		return models.GetTodoResponse{}, fmt.Errorf("null struct")
	}

	if todo.ActivityID == 0 {
		return models.GetTodoResponse{}, fmt.Errorf("null activity id")
	}

	todo, err := u.todoRepository.Create(todo)
	if err != nil {
		return models.GetTodoResponse{}, err
	}

	return models.GetTodoResponse{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Update(todo models.Todo, todoID int) (models.GetTodoResponse, error) {
	foundTodo, err := u.todoRepository.FindByID(todoID)
	if err != nil {
		return models.GetTodoResponse{}, err
	}

	if todo.Title != "" {
		foundTodo.Title = todo.Title
	}

	if todo.Priority != "" {
		foundTodo.Priority = todo.Priority
	}

	if todo.IsActive != foundTodo.IsActive {
		foundTodo.IsActive = todo.IsActive
	}

	updatedTodo, err := u.todoRepository.Update(foundTodo, todoID)
	if err != nil {
		return models.GetTodoResponse{}, err
	}

	return models.GetTodoResponse{
		ID:         updatedTodo.ID,
		ActivityID: updatedTodo.ActivityID,
		Title:      updatedTodo.Title,
		IsActive:   updatedTodo.IsActive,
		Priority:   updatedTodo.Priority,
		CreatedAt:  updatedTodo.CreatedAt,
		UpdatedAt:  foundTodo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Delete(todoID int) error {
	if _, err := u.todoRepository.FindByID(todoID); err != nil {
		return err
	}

	return u.todoRepository.Delete(todoID)
}

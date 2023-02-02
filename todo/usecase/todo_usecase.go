package usecase

import (
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

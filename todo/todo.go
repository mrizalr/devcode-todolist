package todo

import "github.com/mrizalr/devcode-todolist/models"

type TodoRepository interface {
	Find(activityID interface{}) ([]models.Todo, error)
}

type TodoUsecase interface {
	GetAll(activityID interface{}) ([]models.GetTodoResponse, error)
}

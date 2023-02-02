package todo

import "github.com/mrizalr/devcode-todolist/models"

type TodoRepository interface {
	Find(activityID interface{}) ([]models.Todo, error)
	FindByID(todoID int) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo, todoID int) (models.Todo, error)
	Delete(todoID int) error
}

type TodoUsecase interface {
	GetAll(activityID interface{}) ([]models.GetTodoResponse, error)
	GetByID(todoID int) (models.GetTodoResponse, error)
	Create(todo models.Todo) (models.GetTodoResponse, error)
	Update(todo models.Todo, todoID int) (models.GetTodoResponse, error)
	Delete(todoID int) error
}

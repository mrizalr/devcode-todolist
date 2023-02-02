package repository

import (
	"github.com/mrizalr/devcode-todolist/models"
	"github.com/mrizalr/devcode-todolist/todo"
	"gorm.io/gorm"
)

type mysqlTodoRepository struct {
	db *gorm.DB
}

func NewMysqlTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &mysqlTodoRepository{db}
}

func (r *mysqlTodoRepository) Find(activityID interface{}) ([]models.Todo, error) {
	todos := []models.Todo{}

	tx := r.db.Debug().Preload("Activity")
	if _, ok := activityID.(int); ok {
		tx.Where("activity_id = ?", activityID)
	}
	tx.Find(&todos)
	return todos, tx.Error
}

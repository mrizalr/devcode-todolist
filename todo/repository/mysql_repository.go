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

	tx := r.db.Preload("Activity")
	if _, ok := activityID.(int); ok {
		tx.Where("activity_group_id = ?", activityID)
	}
	tx.Find(&todos)
	return todos, tx.Error
}

func (r *mysqlTodoRepository) FindByID(todoID int) (models.Todo, error) {
	todo := models.Todo{}
	tx := r.db.Preload("Activity").Where("todo_id = ?", todoID).First(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Create(todo models.Todo) (models.Todo, error) {
	tx := r.db.Create(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Update(todo models.Todo, todoID int) (models.Todo, error) {
	tx := r.db.Where("todo_id = ?", todoID).Updates(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Delete(todoID int) error {
	return r.db.Where("todo_id = ?", todoID).Delete(&models.Todo{}).Error
}

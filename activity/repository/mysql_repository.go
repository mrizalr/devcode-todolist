package repository

import (
	"github.com/mrizalr/devcode-todolist/activity"
	"github.com/mrizalr/devcode-todolist/models"
	"gorm.io/gorm"
)

type mysqlActivityRepository struct {
	db *gorm.DB
}

func NewMysqlActivityRepository(db *gorm.DB) activity.MysqlActivityRepository {
	return &mysqlActivityRepository{db}
}

func (r *mysqlActivityRepository) FindAll() ([]models.Activity, error) {
	activities := []models.Activity{}
	tx := r.db.Find(&activities)
	return activities, tx.Error
}

func (r *mysqlActivityRepository) FindByID(activityID int) (models.Activity, error) {
	activity := models.Activity{}
	tx := r.db.Where("activity_id = ?", activityID).First(&activity)
	return activity, tx.Error
}

func (r *mysqlActivityRepository) Create(activity models.Activity) (models.Activity, error) {
	tx := r.db.Create(&activity)
	return activity, tx.Error
}

func (r *mysqlActivityRepository) Update(activity models.Activity, activityID int) (models.Activity, error) {
	tx := r.db.Where("activity_id = ?", activityID).Updates(&activity)
	return activity, tx.Error
}

func (r *mysqlActivityRepository) Delete(activityID int) error {
	return r.db.Where("activity_id = ?", activityID).Delete(&models.Activity{}).Error
}

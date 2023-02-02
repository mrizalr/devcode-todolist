package activity

import "github.com/mrizalr/devcode-todolist/models"

type MysqlActivityRepository interface {
	FindAll() ([]models.Activity, error)
	FindByID(activityID int) (models.Activity, error)
	Create(activity models.Activity) (models.Activity, error)
	Update(activity models.Activity, activityID int) (models.Activity, error)
	Delete(activityID int) error
}

type ActivityUsecase interface {
	GetAll() ([]models.Activity, error)
	GetByID(activityID int) (models.Activity, error)
	Create(activity models.Activity) (models.Activity, error)
	Update(activity models.Activity, activityID int) (models.Activity, error)
	Delete(activityID int) error
}

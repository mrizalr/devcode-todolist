package usecase

import (
	"fmt"

	"github.com/mrizalr/devcode-todolist/activity"
	"github.com/mrizalr/devcode-todolist/models"
)

type activityUsecase struct {
	activityRepository activity.MysqlActivityRepository
}

func NewActivityUsecase(activityRepository activity.MysqlActivityRepository) activity.ActivityUsecase {
	return &activityUsecase{activityRepository}
}

func (u *activityUsecase) GetAll() ([]models.Activity, error) {
	return u.activityRepository.FindAll()
}

func (u *activityUsecase) GetByID(activityID int) (models.Activity, error) {
	return u.activityRepository.FindByID(activityID)
}

func (u *activityUsecase) Create(activity models.Activity) (models.Activity, error) {
	if activity.Title == "" {
		return models.Activity{}, fmt.Errorf("null title")
	}

	return u.activityRepository.Create(activity)
}

func (u *activityUsecase) Update(activity models.Activity, activityID int) (models.Activity, error) {
	if (activity == models.Activity{}) {
		return models.Activity{}, fmt.Errorf("null struct")
	}

	foundActivity, err := u.activityRepository.FindByID(activityID)
	if err != nil {
		return models.Activity{}, err
	}

	if activity.Title != "" {
		foundActivity.Title = activity.Title
	}

	if activity.Email != "" {
		foundActivity.Email = activity.Email
	}

	return u.activityRepository.Update(foundActivity, activityID)
}

func (u *activityUsecase) Delete(activityID int) error {
	if _, err := u.activityRepository.FindByID(activityID); err != nil {
		return err
	}

	return u.activityRepository.Delete(activityID)
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	activityHttp "github.com/mrizalr/devcode-todolist/activity/delivery/http"
	activityRepository "github.com/mrizalr/devcode-todolist/activity/repository"
	activityUsecase "github.com/mrizalr/devcode-todolist/activity/usecase"
	"gorm.io/gorm"
)

func MapHandler(app *fiber.App, db *gorm.DB) {
	activityRepo := activityRepository.NewMysqlActivityRepository(db)

	activityUcase := activityUsecase.NewActivityUsecase(activityRepo)

	activityHandler := activityHttp.NewActivityHandler(activityUcase)

	activityHttp.MapRoutes(app, *activityHandler)
}

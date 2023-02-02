package handlers

import (
	"github.com/gofiber/fiber/v2"
	activityHttp "github.com/mrizalr/devcode-todolist/activity/delivery/http"
	activityRepository "github.com/mrizalr/devcode-todolist/activity/repository"
	activityUsecase "github.com/mrizalr/devcode-todolist/activity/usecase"
	todoHttp "github.com/mrizalr/devcode-todolist/todo/delivery/http"
	todoRepository "github.com/mrizalr/devcode-todolist/todo/repository"
	todoUsecase "github.com/mrizalr/devcode-todolist/todo/usecase"
	"gorm.io/gorm"
)

func MapHandler(app *fiber.App, db *gorm.DB) {
	activityRepo := activityRepository.NewMysqlActivityRepository(db)
	todoRepo := todoRepository.NewMysqlTodoRepository(db)

	activityUcase := activityUsecase.NewActivityUsecase(activityRepo)
	todoUcase := todoUsecase.NewTodoUsecase(todoRepo)

	activityHandler := activityHttp.NewActivityHandler(activityUcase)
	todoHandler := todoHttp.NewTodoHandler(todoUcase)

	activityHttp.MapRoutes(app, *activityHandler)
	todoHttp.MapRoutes(app, *todoHandler)
}

package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/devcode-todolist/models"
	"github.com/mrizalr/devcode-todolist/todo"
)

type todoHandler struct {
	todoUsecase todo.TodoUsecase
}

func NewTodoHandler(todoUsecase todo.TodoUsecase) *todoHandler {
	return &todoHandler{todoUsecase}
}

func (h *todoHandler) GetAllTodos() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID, err := strconv.Atoi(c.Query("activity_group_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Invalid type of activity ID",
				})
		}

		todos, err := h.todoUsecase.GetAll(activityID)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).
				JSON(models.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(models.Response{
				Status:  "Success",
				Message: "Success",
				Data:    todos,
			})
	}
}

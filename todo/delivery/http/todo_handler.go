package http

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/devcode-todolist/models"
	"github.com/mrizalr/devcode-todolist/todo"
	"gorm.io/gorm"
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

func (h *todoHandler) GetTodoByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Invalid type of todo ID",
				})
		}

		todo, err := h.todoUsecase.GetByID(todoID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(models.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Todo with ID %d Not Found", todoID),
					})
			}

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
				Data:    todo,
			})
	}
}

func (h *todoHandler) CreateTodo() fiber.Handler {
	type CreateTodoRequest struct {
		Title           string `json:"title"`
		ActivityGroupID int    `json:"activity_group_id"`
		IsActive        bool   `json:"is_active"`
	}
	return func(c *fiber.Ctx) error {
		request := &CreateTodoRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Your request could not be completed as it contains errors. Please check and try again",
				})
		}

		todo, err := h.todoUsecase.Create(models.Todo{
			ActivityID: uint(request.ActivityGroupID),
			Title:      request.Title,
			IsActive:   request.IsActive,
		})
		if err != nil {
			if err.Error() == "null struct" {
				return c.Status(fiber.StatusBadRequest).
					JSON(models.Response{
						Status:  "Bad Request",
						Message: "Title cannot be null",
					})
			}

			return c.Status(fiber.StatusBadGateway).
				JSON(models.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusCreated).
			JSON(models.Response{
				Status:  "Success",
				Message: "Success",
				Data:    todo,
			})
	}
}

func (h *todoHandler) UpdateTodo() fiber.Handler {
	type UpdateTodoRequest struct {
		Title    string `json:"title"`
		Priority string `json:"priority"`
		IsActive bool   `json:"is_active"`
	}
	return func(c *fiber.Ctx) error {
		todoID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Invalid type of todo ID",
				})
		}

		request := &UpdateTodoRequest{}
		err = c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Your request could not be completed as it contains errors. Please check and try again",
				})
		}

		todo, err := h.todoUsecase.Update(models.Todo{
			Title:    request.Title,
			IsActive: request.IsActive,
			Priority: request.Priority,
		}, todoID)
		if err != nil {
			if err.Error() == "null struct" {
				return c.Status(fiber.StatusBadRequest).
					JSON(models.Response{
						Status:  "Bad Request",
						Message: "Title cannot be null",
					})
			}

			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(models.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Todo with ID %d Not Found", todoID),
					})
			}

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
				Data:    todo,
			})
	}
}

func (h *todoHandler) DeleteTodo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.Response{
					Status:  "Bad Request",
					Message: "Invalid type of todo ID",
				})
		}

		err = h.todoUsecase.Delete(todoID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(models.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Todo with ID %d Not Found", todoID),
					})
			}

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
				Data:    struct{}{},
			})
	}
}

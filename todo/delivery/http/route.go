package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(r fiber.Router, h todoHandler) {
	r.Get("/todo-items", h.GetAllTodos())
}

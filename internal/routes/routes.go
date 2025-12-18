package routes

import (
	"go-users-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, h *handler.UserHandler) {
	users := app.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/", h.ListUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

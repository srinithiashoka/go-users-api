package handler

import (
	"strconv"
	"time"

	"go-users-api/internal/models"
	"go-users-api/internal/repository"
	"go-users-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(r *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: r}
}

// ---------- Create ----------
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(400, "Invalid DOB format (YYYY-MM-DD)")
	}

	user := models.User{Name: req.Name, Dob: dob}

	id, err := h.Repo.CreateUser(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.ID = int(id)
	user.Age = service.CalculateAge(user.Dob)

	return c.Status(201).JSON(user)
}

// ---------- Get ----------
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	user.Age = service.CalculateAge(user.Dob)
	return c.JSON(user)
}

// ---------- List ----------
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Repo.ListUsers()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	for i := range users {
		users[i].Age = service.CalculateAge(users[i].Dob)
	}

	return c.JSON(users)
}

// ---------- Update ----------
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(400, "Invalid DOB format")
	}

	user := models.User{ID: id, Name: req.Name, Dob: dob}

	if err := h.Repo.UpdateUser(user); err != nil {
		return fiber.ErrInternalServerError
	}

	user.Age = service.CalculateAge(user.Dob)
	return c.JSON(user)
}

// ---------- Delete ----------
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.Repo.DeleteUser(id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(204)
}

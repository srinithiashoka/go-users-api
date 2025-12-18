package main

import (
	"go-users-api/config"
	"go-users-api/internal/handler"
	"go-users-api/internal/logger"
	"go-users-api/internal/repository"
	"go-users-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Sync()

	db := config.ConnectDB()
	defer db.Close()

	logger.Info("Database connected")

	app := fiber.New()

	repo := repository.NewUserRepository(db)
	handler := handler.NewUserHandler(repo)

	routes.Setup(app, handler)

	logger.Info("Server started", zap.String("port", "8080"))
	app.Listen(":8080")
}

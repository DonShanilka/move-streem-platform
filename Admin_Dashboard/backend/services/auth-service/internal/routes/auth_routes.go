package routes

import (
    "github.com/DonShanilka/auth-service/internal/handlers"
    "github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
    api := app.Group("/auth")
    api.Post("/register", handlers.Register)
    api.Post("/login", handlers.Login)
}

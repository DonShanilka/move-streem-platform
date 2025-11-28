package main

import (
    "log"
    "auth-service/internal/config"
    "auth-service/internal/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Database connect
    config.ConnectDB()

    // Load routes
    routes.SetupAuthRoutes(app)

    log.Fatal(app.Listen(":9002"))
}

package main

import (
	"log"

	"github.com/DonShanilka/auth-service/internal/database"
	"github.com/DonShanilka/auth-service/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect("mongodb://localhost:27017") // <-- Your MongoDB URI

	app := fiber.New()
	routes.SetupAuthRoutes(app)

	log.Println("Auth Service running on :9002")
	log.Fatal(app.Listen(":9002"))
}

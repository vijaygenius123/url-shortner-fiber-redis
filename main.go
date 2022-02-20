package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"url-shortner-fiber-redis/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1", routes.ShortenUrl)
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println()
	}
	app := fiber.New()

	app.Use(logger.New())

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}

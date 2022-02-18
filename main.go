package main

import "github.com/gofiber/fiber"

func setupRoutes(app *fiber.App) {
}

func main() {
	app := fiber.New()

	app.Listen(8080)
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + port)
}

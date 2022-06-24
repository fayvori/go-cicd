package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!! I'm building and serving from pipeline. It's very cool!")
	})

	app.Listen(":" + port)
}

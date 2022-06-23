package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	app := fiber.New()
	route := "/"
	req := httptest.NewRequest("GET", route, nil)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	resp, _ := app.Test(req, 1)
	bytes, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "Hello, World 👋!", string(bytes))
}

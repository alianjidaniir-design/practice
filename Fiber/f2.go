package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	})
	app.Listen(":3000")
}

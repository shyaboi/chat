package main

import (
	"github.com/gofiber/fiber/v2"
)

var chatLog []string

func main() {
	app := fiber.New()

	app.Post("/msg", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		body := c.Body() // user=john
		chatLog = append(chatLog, string(body))
		return c.JSON(chatLog)
	})

	app.Get("/lm", func(c *fiber.Ctx) error {
		// benix := fmt.Sprint(s)
		return c.JSON(chatLog)
	})

	app.Listen(":3000")
}

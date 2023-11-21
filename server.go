package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/messages", func(c *fiber.Ctx) error {
		messages := GetConversationMessages("a")

		return c.Render("results", fiber.Map{
			"Results": messages,
		})
	})

	app.Get("api/messages", func(c *fiber.Ctx) error {
		messages := GetConversationMessages("a")

		return c.JSON(&fiber.Map{
			"Status":  true,
			"Results": messages,
		})
	})

	app.Listen(":3000")
}

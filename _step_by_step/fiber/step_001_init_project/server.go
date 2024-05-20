package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello world!")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { // c for context
		return c.JSON("hello fiber")
	})

	app.Listen(":8080")

}

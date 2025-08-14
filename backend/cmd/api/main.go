package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("app is running..ß.")

	app := fiber.New()

	todoArr := []Todo{}

	app.Use(logger.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		fmt.Println("pinging...")
		return c.JSON(fiber.Map{"message": "Server is healthy"})
	})

	app.Post("/api/create-todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todoArr) + 1
		todoArr = append(todoArr, *todo)
		return c.Status(201).JSON(todo)

	})

	log.Fatal(app.Listen(":8000"))
}

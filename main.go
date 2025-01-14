package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Create a ToDo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}

		todo.ID = strconv.Itoa(len(todos) + 1)
		todos = append(todos, *todo)

		return c.Status(fiber.StatusCreated).JSON(todos)
	})

	// Update a ToDo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	})

	// Delete a ToDo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(todos)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	})
	app.Listen(":4001")
}

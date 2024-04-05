package main

import "github.com/gofiber/fiber/v2"

func main() {

	// Create a new instance of the Fiber app.
	app := fiber.New()

	// Routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	// Routes for the puzzles
	app.Get("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Get("/puzzles/:id", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Post("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Patch("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Delete("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	// Start the server and listen on port 3000
	app.Listen(":3000")
}

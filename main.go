package main

import (
	"github.com/BloomGameStudio/PuzzleService/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    _, err := database.DBInit()
    if err != nil {
        // handle error
    }

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("Pong!")
    })

    app.Listen(":3000")
}
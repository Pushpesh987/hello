package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Define a route handler for the root path
    app.Get("/", func(c *fiber.Ctx) error {
        // Return "Hello, World!" as the response
        return c.SendString("Hello, World!")
    })

    // Start the Fiber app on port 3000
    app.Listen(":3000")
}

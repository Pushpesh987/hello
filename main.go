package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file if it exists
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Create a new Fiber instance
    app := fiber.New()

    // Define a route handler for the root path
    app.Get("/", func(c *fiber.Ctx) error {
        // Return "Hello, World!" as the response
        return c.SendString("Hello, World!")
    })

    // Fetch the port number from the environment variable, default to 3000 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    // Start the Fiber app on the specified port
    log.Fatal(app.Listen(":" + port))
}

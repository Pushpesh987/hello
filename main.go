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
        log.Println("No .env file found, relying on environment variables")
    }

    // Create a new Fiber instance
    app := fiber.New()

    // Define a route handler for the root path
    app.Get("/", func(c *fiber.Ctx) error {
        // Return "Hello, World!" as the response
        return c.SendString("Hello, World!")
    })

    // Fetch the port number from the environment variable, default to 4000 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "4000"
    }

    // Start the Fiber app on the specified port
    log.Printf("Starting server on port %s...", port)
    log.Fatal(app.Listen(":" + port))
}

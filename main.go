package main

import (
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,OPTIONS",
        AllowHeaders: "Content-Type",
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok", "message": "Maffin API 🐱"})
    })

    app.Post("/api/chat", func(c *fiber.Ctx) error {
        // handler ของคุณ
        return nil
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    app.Listen(":" + port)
}
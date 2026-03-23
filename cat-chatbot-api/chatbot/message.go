package chatbot

import (
    "strings"
    "github.com/gofiber/fiber/v2"
)

type ChatRequest struct {
    Message string `json:"message"`
    UserId  string `json:"user_id"`
}

func ChatHandler(c *fiber.Ctx) error {
    var req ChatRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
    }

    if strings.TrimSpace(req.Message) == "" {
        return c.Status(400).JSON(fiber.Map{"error": "Message is required"})
    }

    resp := GetBotReply(req.Message)
    return c.JSON(resp)
}
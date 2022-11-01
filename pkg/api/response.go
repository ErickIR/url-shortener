package api

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RespondError(c *fiber.Ctx, statusCode int, body interface{}) error {
	c.Set("Content-type", "application/json")

	data, err := json.Marshal(body)
	if err != nil {
		log.Println("encoding error: ", err)
	}

	return fiber.NewError(statusCode, string(data))
}

func RespondWithJSON(c *fiber.Ctx, statusCode int, body interface{}) error {
	c.Set("Content-type", "application/json")

	data, err := json.Marshal(body)
	if err != nil {
		log.Println("encoding error: ", err)
	}

	return c.SendString(string(data))
}

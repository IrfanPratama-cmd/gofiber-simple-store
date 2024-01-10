package lib

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetXUserID(c *fiber.Ctx) *uuid.UUID {
	userID := c.Locals("userID")
	id := userID.(string)
	if id != "" {
		if current, err := uuid.Parse(id); nil == err {
			return &current
		}
	}

	return nil
}

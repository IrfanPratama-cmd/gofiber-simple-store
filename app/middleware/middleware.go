package middleware

import (
	"test-api/app/utility"

	"github.com/gofiber/fiber/v2"
)

func Middleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// _, err := utility.VerifyToken(token)
	claims, err := utility.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// role := claims["role"].(string)
	// if role != "user" {
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"message": "forbidden access",
	// 	})
	// }

	userID := claims["user_id"].(string)

	c.Locals("userInfo", claims)
	c.Locals("userID", userID)

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}

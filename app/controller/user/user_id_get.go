package user

import (
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	db := services.DB

	var user model.User
	result := db.Model(&model.User{}).Where(`id = ?`, id).First(&user)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	return c.JSON(user)
}

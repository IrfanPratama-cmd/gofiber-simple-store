package user

import (
	"fmt"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	var user model.User
	result := db.Model(&user).Where(`id = ?`, id).First(&user)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	result.Delete(&user)

	message := fmt.Sprintf(`User with id %s has been deleted`, id)
	return c.JSON(fiber.Map{
		"message": message,
	})
}

package user

import (
	"fmt"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func PutUser(c *fiber.Ctx) error {
	var UserAPI model.UserAPI
	if err := c.BodyParser(&UserAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var user model.User

	if rowsAffected := db.First(&user, `id = ?`, id).RowsAffected; rowsAffected < 1 {
		return c.Status(404).JSON(fiber.Map{
			"message": "ID Not Found",
		})
	}

	update := &model.User{UserAPI: UserAPI}
	db.Model(&model.User{}).Where(`id = ?`, id).Updates(update)
	message := fmt.Sprintf(`User with id %s has been updated`, id)

	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    update,
	})

}

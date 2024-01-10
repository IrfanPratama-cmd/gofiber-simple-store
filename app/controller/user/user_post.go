package user

import (
	"test-api/app/model"
	"test-api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {
	var userAPI model.UserAPI

	db := services.DB

	if err := c.BodyParser(&userAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(userAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var email model.User
	checkEmail := db.Model(&model.User{}).Where(`email = ?`, userAPI.Email).First(&email)

	if checkEmail.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is already used",
		})
	}

	user := &model.User{UserAPI: userAPI}
	db.Model(&model.User{}).Create(user)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

package auth

import (
	"log"
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(model.UserAccountAPI)
	db := services.DB

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var userAccount model.UserAccount
	err := db.First(&userAccount, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	checkPassword := lib.CheckPassword(loginRequest.Password, userAccount.Password)
	if !checkPassword {
		return c.Status(404).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["user_id"] = userAccount.ID
	claims["email"] = userAccount.Email
	claims["exp"] = time.Now().Add(time.Minute * 10000).Unix()

	token, errGenerateToken := lib.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

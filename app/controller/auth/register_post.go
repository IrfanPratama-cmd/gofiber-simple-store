package auth

import (
	"test-api/app/model"
	"test-api/app/services"
	"test-api/app/utility"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	register := new(model.UserAccountAPI)

	db := services.DB

	if err := c.BodyParser(&register); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(register)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var email model.UserAccount
	checkEmail := db.Model(&model.User{}).Where(`email = ?`, register.Email).First(&email)

	if checkEmail.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is already used",
		})
	}

	hash, _ := utility.HashPassword(register.Password)

	roleName := "User"

	var role model.Role
	db.Model(&model.Role{}).Where(`role_name = ?`, roleName).First(&role)

	var account model.UserAccount
	account.Email = register.Email
	account.Password = hash
	db.Create(&account)

	var user model.User
	user.Email = register.Email
	user.UserAccountID = account.ID
	db.Create(&user)

	var userRole model.RoleUser
	userRole.UserID = account.ID
	userRole.RoleID = role.ID
	db.Create(&userRole)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    account,
		"role":    userRole,
	})
}

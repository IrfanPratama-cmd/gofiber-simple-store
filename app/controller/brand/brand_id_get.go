package brand

import (
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetBrandID(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var brand model.Brand
	result := db.Model(&brand).Where("id = ?", id).First(&brand)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand not found",
		})
	}

	return lib.OK(c, brand)
}

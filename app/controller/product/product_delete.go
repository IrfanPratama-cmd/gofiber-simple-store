package product

import (
	"fmt"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	var product model.Product
	result := db.Model(&product).Where(`id = ?`, id).First(&product)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	result.Delete(&product)

	message := fmt.Sprintf(`Product with id %s has been deleted`, id)
	return c.JSON(fiber.Map{
		"message": message,
	})
}

package product

import (
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetProductID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	userID := lib.GetXUserID(c)

	var product model.Product
	result := db.Model(&product).
		Select(`products.id, products.created_at, products.updated_at,
			products.product_name, products.description, products.price, products.quantity,
			products.thumbnail,
			c.id "Category__id",
			c.created_at "Category__created_at",
			c.updated_at "Category__updated_at",
			c.category_name "Category__category_name",
			c.category_code "Category__category_code",
			u.name "User__name",
			u.created_at "User__created_at",
			u.updated_at "User__updated_at",
			u.email "User__email"`).
		Joins(`LEFT JOIN users u on u.user_account_id = products.user_id`).
		Joins(`LEFT JOIN categories c on c.id = products.category_id`).
		Where(`products.id = ?`, id).
		Where(`products.user_id = ?`, userID).
		First(&product)

	if result.RowsAffected < 1 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product ID Not Found",
		})
	}

	var asset []model.ProductAsset
	db.Where(`product_id = ?`, product.ID).Find(&asset)
	return c.Status(200).JSON(fiber.Map{
		"product":       product,
		"product_asset": asset,
	})
}

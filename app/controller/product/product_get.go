package product

import (
	"log"
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetProduct(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	userID := lib.GetXUserID(c)

	path := c.Path()
	method := c.Method()

	log.Println("Path :", path)
	log.Println("Method :", method)

	var product model.Product

	mod := db.Model(&product).
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
		// Joins(`LEFT JOIN product_assets pa on pa.product_id = products.id`).
		Where(`user_id = ?`, userID)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})
	return c.Status(200).JSON(page)
}

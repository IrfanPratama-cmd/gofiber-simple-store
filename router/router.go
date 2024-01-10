package router

import (
	"test-api/app/controller/auth"
	"test-api/app/controller/brand"
	"test-api/app/controller/category"
	"test-api/app/controller/product"
	"test-api/app/controller/user"
	"test-api/app/middleware"
	"test-api/app/utility"
	"test-api/config"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world"})
	})

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)

	app.Get("/user", middleware.Middleware, user.GetUser)
	app.Post("/user", middleware.Middleware, user.PostUser)
	app.Put("/user/:id", middleware.Middleware, user.PutUser)
	app.Get("/user/:id", middleware.Middleware, user.GetUserID)
	app.Delete("/user/:id", middleware.Middleware, user.DeleteUser)

	app.Get("/product", middleware.Middleware, product.GetProduct)
	app.Post("/product", middleware.Middleware, utility.HandleSingleFile, product.PostProduct)
	app.Get("/product/:id", middleware.Middleware, product.GetProductID)
	app.Put("/product/:id", middleware.Middleware, utility.HandleSingleFile, product.PutProduct)
	app.Delete("/product/:id", middleware.Middleware, product.DeleteProduct)

	app.Get("/category", middleware.Middleware, category.GetCategory)
	app.Post("/category", middleware.Middleware, category.PostCategory)
	app.Get("/category/:id", middleware.Middleware, category.GetCategoryID)
	app.Put("/category/:id", middleware.Middleware, category.PutCategory)
	app.Delete("/category/:id", middleware.Middleware, category.DeleteCategory)

	app.Get("/brands", middleware.Middleware, brand.GetBrand)
	app.Post("/brands", middleware.Middleware, brand.PostBrand)
	app.Get("/brands/:id", middleware.Middleware, brand.GetBrandID)
	app.Put("/brands/:id", middleware.Middleware, brand.PutBrand)
	app.Delete("/brands/:id", middleware.Middleware, brand.DeleteBrand)

	app.Static("/public", config.ProjectRootPath+"/public/asset")
}

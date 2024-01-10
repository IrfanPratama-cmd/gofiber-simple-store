package product

import (
	"fmt"
	"log"
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostProduct(c *fiber.Ctx) error {
	var productAPI model.ProductRequest

	db := services.DB
	userID := lib.GetXUserID(c)

	if err := c.BodyParser(&productAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(productAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Handler File
	fileName := c.Locals("thumbnail")

	if fileName == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "thumbnail is required",
		})
	}

	thumbnail := fmt.Sprintf("%v", fileName)

	// log.Println("File :", fileName)

	var product model.Product
	product.ProductName = productAPI.ProductName
	product.Description = productAPI.Description
	product.Price = productAPI.Price
	product.Quantity = productAPI.Quantity
	product.CategoryID = productAPI.CategoryID
	product.BrandID = productAPI.BrandID
	product.UserID = userID
	product.Thumbnail = thumbnail
	db.Create(&product)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["product_asset"]

	// return c.Status(200).JSON(files)

	var productAsset model.ProductAsset
	for _, file := range files {
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/product/%s", file.Filename))
		filePath := fmt.Sprintf("./public/product/%s", file.Filename)
		if errSaveFile != nil {
			log.Println("Upload Failed ")
		}
		db.Model(&productAsset).Create(map[string]interface{}{
			"id":         uuid.New(),
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"product_id": product.ID,
			"file_name":  file.Filename,
			"file_path":  filePath,
		})
		// Save the file
		// productAsset.ProductID = product.ID
		// productAsset.FileName = file.Filename
		// productAsset.FilePath = filePath
		// if err := db.Create(&productAsset); err != nil {
		// 	log.Println(err)
		// }
	}
	// db.Create(&productAsset)

	// var productAsset model.ProductAsset

	// asset := c.Locals("product_asset")
	// assets := asset.([]string)
	// for _, asset := range assets {
	// 	productAsset.ProductID = product.ID
	// 	productAsset.FileName = asset
	// 	db.Create(&productAsset)
	// 	// if err := db.Create(&productAsset).Error; err != nil {
	// 	// 	return err
	// 	// }
	// }

	// Map Interface Example
	// db.Model(&product).Create(map[string]interface{}{
	// 	"product_name": productAPI.ProductName,
	// 	"description":  productAPI.Description,
	// 	"price":        productAPI.Price,
	// 	"quantity":     productAPI.Quantity,
	// 	"user_id":      userID,
	// })

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})
}

package utility

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	file, errFile := c.FormFile("thumbnail")
	if errFile != nil {
		log.Println("Error file = ", errFile)
	}

	var fileName string

	if file != nil {
		fileName = file.Filename
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/thumbnail/%s", fileName))
		if errSaveFile != nil {
			log.Println("Upload Failed ", errFile)
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	c.Locals("thumbnail", fileName)

	return c.Next()
}

func HandleMultipleFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error read multipart request = ", err)
	}

	files := form.File["asset"]

	var fileNames []string
	for i, file := range files {
		var fileName string

		if file != nil {
			fileName = fmt.Sprintf("%s-%d", file.Filename, i)
			errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/product/%s", fileName))
			if errSaveFile != nil {
				log.Println("Upload Failed ")
			}
		} else {
			log.Println("Nothing file to uploading.")
		}

		if fileName != "" {
			fileNames = append(fileNames, fileName)

		}
		c.Locals("product_asset", fileNames)
	}

	return c.Next()
}

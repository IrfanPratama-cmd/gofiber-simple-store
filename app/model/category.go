package model

import (
	"strings"
	"test-api/app/services"
)

type Category struct {
	Base
	CategoryAPI
}

type CategoryAPI struct {
	CategoryCode *string `json:"category_code,omitempty" validate:"required"`
	CategoryName *string `json:"category_name,omitempty" validate:"required"`
}

func (c Category) Seed() *[]Category {
	data := []Category{}
	items := []string{
		"LP|Laptop",
		"HP|Handphone",
		"TB|Tablet",
		"HD|Hardisk",
		"MN|Monitor",
		"FD|Flashdisk",
		"PC|Computer",
	}

	db := services.DB

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		categoryCode := c[0]
		categoryName := c[1]
		var data Category
		data.CategoryCode = &categoryCode
		data.CategoryName = &categoryName
		db.Create(&data)
	}
	return &data
}

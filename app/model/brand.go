package model

import (
	"strings"
	"test-api/app/services"
)

type Brand struct {
	Base
	BrandAPI
}

type BrandAPI struct {
	BrandCode *string `json:"brand_code,omitempty" validate:"required"`
	BrandName *string `json:"brand_name,omitempty" validate:"required"`
}

func (b Brand) Seed() *[]Brand {
	data := []Brand{}
	items := []string{
		"AS|Asus",
		"SM|Samsung",
		"XM|Xiaomi",
		"IP|Iphone",
		"AC|Acer",
		"OP|Oppo",
		"VO|Vivo",
	}

	db := services.DB

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		brandCode := c[0]
		brandName := c[1]
		var data Brand
		data.BrandCode = &brandCode
		data.BrandName = &brandName
		db.Create(&data)
	}
	return &data
}

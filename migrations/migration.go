package migrations

import (
	"fmt"
	"log"
	"test-api/app/model"
	"test-api/app/services"
)

func RunMigration() {
	err := services.DB.AutoMigrate(
		&model.User{},
		&model.UserAccount{},
		&model.Role{},
		&model.RoleUser{},
		&model.Category{},
		&model.Brand{},
		&model.Product{},
		&model.ProductAsset{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")

}

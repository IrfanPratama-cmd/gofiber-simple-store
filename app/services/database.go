package services

import (
	"fmt"

	"github.com/morkid/paginate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PG *paginate.Pagination

func InitDatabase() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/gofiber-simple-store?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}

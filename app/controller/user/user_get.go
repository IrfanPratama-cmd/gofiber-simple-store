package user

import (
	"test-api/app/lib"
	"test-api/app/model"
	"test-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetUser(c *fiber.Ctx) error {
	var user []model.User
	db := services.DB
	pg := paginate.New()
	mod := db.Model(&model.User{}).Find(&user)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.User{})
	return lib.OK(c, page)
}

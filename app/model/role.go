package model

import (
	"strings"
	"test-api/app/services"
)

type Role struct {
	Base
	RoleAPI
}

type RoleAPI struct {
	RoleName    *string `json:"role_name,omitempty" validate:"required"`
	Description *string `json:"description,omitempty"`
}

func (r Role) Seed() *[]Role {
	data := []Role{}
	items := []string{
		"Admin|Role Admin",
		"User|Role User",
	}

	db := services.DB

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		roleName := c[0]
		description := c[1]
		var data Role
		if db.Where(Role{RoleAPI: RoleAPI{RoleName: &roleName}}).First(&data).Error != nil {
			data.RoleName = &roleName
			data.Description = &description
			db.Create(&data)
		}
	}
	return &data
}

package model

import "github.com/google/uuid"

type User struct {
	Base
	UserAPI
	UserAccount *UserAccount `json:"user_account,omitempty"`
}

type UserAPI struct {
	Name          string     `json:"name,omitempty" validate:"required"`
	Email         string     `json:"email,omitempty" validate:"required"`
	Phone         string     `json:"phone,omitempty"`
	Address       string     `json:"address,omitempty"`
	UserAccountID *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
}

type UserDetail struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
}

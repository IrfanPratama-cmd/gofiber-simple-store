package model

import "github.com/google/uuid"

type RoleUser struct {
	Base
	RoleUserAPI
	User *UserAccount `json:"user,omitempty"`
	Role *Role        `json:"role,omitempty"`
}

type RoleUserAPI struct {
	UserID *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	RoleID *uuid.UUID `json:"role_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
}

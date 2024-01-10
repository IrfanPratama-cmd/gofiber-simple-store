package model

import "github.com/google/uuid"

type UserAsset struct {
	Base
	UserAssetAPI
}

type UserAssetAPI struct {
	UserID   *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	FileName string     `json:"file_name,omitempty" validate:"required"`
}

package models

import (
	"time"

	"final-project/helpers"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"type:varchar(255)" validate:"required" json:"name"`
	SocialMediaUrl string     `gorm:"type:varchar(255)" validate:"required" json:"social_media_url"`
	UserID         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	errCreate := helpers.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}

	return nil
}

func (sm *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	errUpdate := helpers.ValidateStruct(sm)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	return nil
}

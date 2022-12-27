package models

import (
	"time"

	"final-project/helpers"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:varchar(255)" validate:"required" json:"title"`
	Caption   string     `gorm:"type:varchar(255)" validate:"-" json:"caption"`
	PhotoUrl  string     `gorm:"type:varchar(255)" validate:"required" json:"photo_url"`
	UserID    uint       `json:"user_id"`
	Comments  []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User      `json:"User,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	errCreate := helpers.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	errUpdate := helpers.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	return nil
}

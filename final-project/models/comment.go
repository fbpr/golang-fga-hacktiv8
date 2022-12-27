package models

import (
	"fmt"
	"time"

	"final-project/helpers"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Message   string     `gorm:"type:text" validate:"required" json:"message"`
	UserID    uint       `json:"user_id"`
	PhotoID   uint       `json:"photo_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User
	Photo     *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	errCreate := helpers.ValidateStruct(c)

	fmt.Println(c.Message)
	if errCreate != nil {
		err = errCreate
		return
	}
	fmt.Println(c.Message)
	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	errUpdate := helpers.ValidateStruct(c)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	return nil
}

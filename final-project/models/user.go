package models

import (
	"time"

	"final-project/helpers"

	"gorm.io/gorm"
)

type User struct {
	ID          uint        `gorm:"primaryKey" json:"id,omitempty"`
	Email       string      `gorm:"type:varchar(255);uniqueIndex;not null" validate:"required,email" json:"email,omitempty"`
	Username    string      `gorm:"type:varchar(255);uniqueIndex;not null" validate:"required" json:"username,omitempty"`
	Password    string      `gorm:"type:varchar(255);not null" validate:"required,min=6" json:"password,omitempty"`
	Age         uint8       `gorm:"type:integer" validate:"required,min=8" json:"age,omitempty"`
	Photos      []Photo     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Comments    []Comment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	SocialMedia SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-" json:"-"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	errCreate := helpers.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	errUpdate := helpers.ValidateStruct(u)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	return nil
}

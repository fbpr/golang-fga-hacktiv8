package models

import (
	"go-jwt/helpers"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `gorm:"not null" json:"full_name" form:"full_name" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := helpers.TranslateError(validate, u); 
	if errCreate != nil {
		err = errCreate
        return
    }

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

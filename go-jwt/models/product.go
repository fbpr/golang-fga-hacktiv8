package models

import (
	"go-jwt/helpers"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	UserID uint
	User *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := helpers.TranslateError(validate, p); 
	if errCreate != nil {
		err = errCreate
        return
    }

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errUpdate := helpers.TranslateError(validate, p); 
	if errUpdate != nil {
		err = errUpdate
        return
    }


	err = nil
	return
}
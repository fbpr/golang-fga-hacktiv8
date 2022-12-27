package models

import (
	_ "gorm.io/gorm"
)

type Item struct {
	ItemID uint `gorm:"not null;primaryKey" json:"lineItemId"`
	ItemCode string `gorm:"type:varchar(255)" json:"itemCode"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Quantity uint `gorm:"type:integer" json:"quantity"`
	OrderID uint `json:"order_id"`
}
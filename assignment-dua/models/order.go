package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Order struct {
	OrderID uint `gorm:"not null;primaryKey" json:"orderId"`
	CustomerName string `gorm:"type:varchar(255)" json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
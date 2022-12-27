package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"assignment-dua/database"
	"assignment-dua/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder (c *gin.Context) {
	db := database.GetDB()

	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	order.OrderedAt = time.Now()
	err := db.Debug().Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order success"})
}

func GetOrder (c *gin.Context) {
	db := database.GetDB()

	order := []models.Order{}

	if err := db.Debug().Preload("Items").Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func UpdateOrderById (c *gin.Context) {
	db := database.GetDB()

	order := models.Order{}
	err := db.Debug().Preload("Items").First(&order, c.Param("orderId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	newOrder := models.Order{OrderID: order.OrderID}

	err = json.NewDecoder(c.Request.Body).Decode(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(newOrder).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func DeleteOrderById (c *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	err := db.First(&order, c.Param("orderId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Debug().Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
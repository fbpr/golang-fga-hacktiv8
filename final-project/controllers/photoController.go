package controllers

import (
	"net/http"

	"final-project/database"
	"final-project/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AddPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	photo.UserID = uint(userData["id"].(float64))

	if err := db.Debug().Create(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add a photo;" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()

	var photos []models.Photo
	if err := db.Debug().Joins("User", db.Select("email", "username")).Find(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, photos)
}

func UpdatePhotoByID(c *gin.Context) {
	db := database.GetDB()

	var photo, newPhoto models.Photo
	if err := db.First(&photo, c.Param("photoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "photo not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	err := db.Model(&photo).Updates(newPhoto).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}

func DeletePhotoByID(c *gin.Context) {
	db := database.GetDB()

	err := db.Delete(&models.Photo{}, c.Param("photoId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})
}

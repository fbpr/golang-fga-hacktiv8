package controllers

import (
	"net/http"

	"final-project/database"
	"final-project/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AddSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var socialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	if err := db.Debug().Create(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add a socialMedia;" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()

	var socialMedias []models.SocialMedia
	if err := db.Debug().Joins("User", db.Select("id", "username")).Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, socialMedias)
}

func UpdateSocialMediaByID(c *gin.Context) {
	db := database.GetDB()

	var socialMedia, newSocialMedia models.SocialMedia
	if err := db.First(&socialMedia, c.Param("socialMediaId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "social media not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	err := db.Model(&socialMedia).Updates(newSocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func DeleteSocialMediaByID(c *gin.Context) {
	db := database.GetDB()

	err := db.Delete(&models.SocialMedia{}, c.Param("socialMediaId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "social media not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}

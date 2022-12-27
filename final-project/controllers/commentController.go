package controllers

import (
	"net/http"

	"final-project/database"
	"final-project/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AddComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	comment.UserID = uint(userData["id"].(float64))

	if err := db.Debug().Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add a comment, photo not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func GetComments(c *gin.Context) {
	db := database.GetDB()

	var comments []models.Comment
	if err := db.Debug().Joins("User", db.Select("id", "email", "username")).Joins("Photo", db.Select("id", "title", "caption", "photo_url", "user_id")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func UpdateCommentByID(c *gin.Context) {
	db := database.GetDB()

	var newComment models.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	commentId := c.Param("commentId")
	err := db.Debug().Model(&newComment).Take(&newComment, commentId).Where("id = ?", commentId).Updates(&newComment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         newComment.ID,
		"message":    newComment.Message,
		"photo_id":   newComment.PhotoID,
		"user_id":    newComment.UserID,
		"updated_at": newComment.UpdatedAt,
	})
}

func DeleteCommentByID(c *gin.Context) {
	db := database.GetDB()

	err := db.Delete(&models.Comment{}, c.Param("commentId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}

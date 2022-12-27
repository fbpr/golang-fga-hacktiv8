package controllers

import (
	"net/http"

	"final-project/database"
	"final-project/helpers"
	"final-project/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDB()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	if err := db.Debug().Where("email = ? OR username = ?", user.Email, user.Username).FirstOrCreate(&user); err.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or email already taken",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func LoginUser(c *gin.Context) {
	db := database.GetDB()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	inputPassword := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid email",
		})

		return
	}

	comparePass := helpers.VerifyPassword(user.Password, inputPassword)
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid password",
		})

		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUserByID(c *gin.Context) {
	db := database.GetDB()

	var user, newData models.User
	if err := db.First(&user, c.Param("userId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})

		return
	}

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})

		return
	}

	err := db.Model(&user).Updates(newData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func DeleteUserByID(c *gin.Context) {
	db := database.GetDB()

	err := db.Delete(&models.User{}, c.Param("userId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})
}

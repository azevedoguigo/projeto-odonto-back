package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Create(&user)

	userResponse := models.UserResponse{
		Model: user.Model,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, userResponse)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found!",
		})
	}

	userResponse := models.UserResponse{
		Model: user.Model,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, userResponse)
}

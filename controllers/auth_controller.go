package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/azevedoguigo/projeto-odonto-back.git/utils"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email!",
		})
		return
	}

	if !user.CheckPassword(credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password!",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

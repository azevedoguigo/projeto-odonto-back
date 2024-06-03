package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/gin-gonic/gin"
)

func CreateDentist(ctx *gin.Context) {
	var dentist models.Dentist

	if err := ctx.ShouldBindJSON(&dentist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Create(&dentist)

	ctx.JSON(http.StatusCreated, dentist)
}

func GetDentistById(ctx *gin.Context) {
	id := ctx.Param("id")

	var dentist models.Dentist

	if err := config.DB.First(&dentist, id).Error; err != nil {
		ctx.JSON(http.StatusNotExtended, gin.H{
			"error": "Dentist not found!",
		})
	}

	ctx.JSON(http.StatusOK, dentist)
}

func GetDentists(ctx *gin.Context) {
	var dentists []models.Dentist

	if err := config.DB.Find(&dentists).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, dentists)
}

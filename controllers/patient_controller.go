package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/gin-gonic/gin"
)

func CreatePatient(ctx *gin.Context) {
	var patient models.Patient

	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Create(&patient)
	ctx.JSON(http.StatusCreated, patient)
}

func GetPatientById(ctx *gin.Context) {
	id := ctx.Param("id")

	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found!",
		})
	}

	ctx.JSON(http.StatusOK, patient)
}

func GetPatients(ctx *gin.Context) {
	var patients []models.Patient

	if err := config.DB.Find(&patients).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, patients)
}

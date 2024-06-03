package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/gin-gonic/gin"
)

type CreateServiceInput struct {
	DentistId uint `json:"dentist-id"`
	PatientId uint `json:"patient-id"`
}

func CreateService(ctx *gin.Context) {
	var input CreateServiceInput
	var dentist models.Dentist
	var patient models.Patient

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := config.DB.First(&dentist, input.DentistId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Dentist not found!",
		})
	}

	if err := config.DB.First(&patient, input.PatientId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found!",
		})
	}

	service := models.Service{
		DentistId: input.DentistId,
		PatientId: input.PatientId,
	}

	config.DB.Create(&service)

	ctx.JSON(http.StatusCreated, service)
}

func GetServiceById(ctx *gin.Context) {
	id := ctx.Param("id")

	var service models.Service

	if err := config.DB.First(&service, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found!",
		})
	}

	ctx.JSON(http.StatusOK, service)
}

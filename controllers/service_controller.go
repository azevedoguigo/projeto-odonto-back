package controllers

import (
	"net/http"

	"github.com/azevedoguigo/projeto-odonto-back.git/config"
	"github.com/azevedoguigo/projeto-odonto-back.git/models"
	"github.com/gin-gonic/gin"
)

type CreateOrUpdateServiceInput struct {
	DentistId uint `json:"dentist-id"`
	PatientId uint `json:"patient-id"`
}

func CreateService(ctx *gin.Context) {
	var input CreateOrUpdateServiceInput
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

func GetServices(ctx *gin.Context) {
	var services []models.Service

	if err := config.DB.Find(&services).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, services)
}

func UpdateService(ctx *gin.Context) {
	id := ctx.Param("id")

	var service models.Service

	if err := config.DB.First(&service, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found!",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Save(&service)

	ctx.JSON(http.StatusOK, service)
}

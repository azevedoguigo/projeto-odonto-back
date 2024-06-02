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

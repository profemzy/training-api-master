package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infotitanz/dancerapy-api/internal/helper"
	"github.com/infotitanz/dancerapy-api/internal/model"
)

func AddTraining(context *gin.Context) {
	var input model.Training
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = user.ID
	savedTraining, err := input.CreateTraining()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedTraining})
}

func UpdateTraining(context *gin.Context) {
	id := context.Param("id")
	trainingID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var training model.Training

	if err := context.ShouldBindJSON(&training); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedTraining, err := model.UpdateTraining(uint(trainingID), training)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": savedTraining})
}

func GetUserTrainings(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user.Trainings})
}

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "API is alive"})
}

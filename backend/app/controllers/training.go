package controllers

import (
	"grass_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingController struct{}

var trainingModel = new(models.Training)

func (t TrainingController) Retrieve(c *gin.Context) {
	if c.Param("date") != "" {
		training, err := trainingModel.GetByDate(c.Param("date"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve training", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"training": training})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

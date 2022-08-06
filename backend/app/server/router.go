package server

import (
	"grass_backend/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("v1")
	{
		trainingGroup := v1.Group("training")
		{
			training := new(controllers.TrainingController)
			trainingGroup.GET("/", training.RetrieveList)
			trainingGroup.GET("/:date", training.Retrieve)
		}
	}
	return router
}

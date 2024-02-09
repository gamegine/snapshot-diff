package routes

import (
	"snapshot-diff/controllers"
	"snapshot-diff/models"

	"github.com/gin-gonic/gin"
)

func VolumesRoute(router *gin.Engine) {
	controllers.Volumes, _ = models.LoadVolumes()
	router.GET("/volumes", controllers.GetVolumes)
	router.GET("/volumes/:volume", controllers.GetVolume)
	router.GET("/volumes/:volume/:snapshot", controllers.GetSnapshot)
}

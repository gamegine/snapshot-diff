package routes

import (
	"snapshot-diff/controllers"

	"github.com/gin-gonic/gin"
)

func VolumesRoute(router *gin.Engine) {
	controllers.InitVolumes()
	router.GET("/volumes", controllers.GetVolumes)
	router.GET("/volumes/:volume", controllers.GetVolume)
	router.GET("/volumes/:volume/:snapshot", controllers.GetSnapshot)
}

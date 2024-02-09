package main

import (
	"snapshot-diff/config"
	"snapshot-diff/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Config()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.VolumesRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

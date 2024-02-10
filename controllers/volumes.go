package controllers

import (
	"snapshot-diff/models"
	"snapshot-diff/utils"

	"github.com/gin-gonic/gin"
)

var Volumes models.Volumes

func GetVolumes(c *gin.Context) {
	c.JSON(200, gin.H{
		"volumes": utils.MapKeys(Volumes),
	})
}

func GetVolume(c *gin.Context) {
	volume := c.Param("volume")
	{
		v := Volumes[volume]
		v.UpdateSnapshotsList()
		Volumes[volume] = v
	}
	c.JSON(200, gin.H{
		"SnapshotsPath": Volumes[volume].SnapshotsPath,
		"Snapshots":     utils.MapKeys(Volumes[volume].Snapshots),
	})
}

func GetSnapshot(c *gin.Context) {
	volume := c.Param("volume")
	snapshot := c.Param("snapshot")
	{
		s := Volumes[volume].Snapshots[snapshot]
		s.LoadFiles()
		Volumes[volume].Snapshots[snapshot] = s
	}
	c.JSON(200, Volumes[volume].Snapshots[snapshot])
}

package controllers

import (
	"fmt"
	"snapshot-diff/models"
	"snapshot-diff/utils"

	"github.com/gin-gonic/gin"
)

func GetVolumes(c *gin.Context) {
	volumes, err := models.LoadVolumes()
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": fmt.Sprintf("LoadVolumes error: %v", err)})
		return
	}
	c.JSON(200, gin.H{
		"volumes": utils.MapKeys(volumes),
	})
}

func GetVolume(c *gin.Context) {
	volume := c.Param("volume")
	v := models.GetVolume(volume)
	snapshots, err := v.GetSnapshots()
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": fmt.Sprintf("UpdateSnapshotsList error: %v", err)})
		return
	}
	c.JSON(200, gin.H{
		"SnapshotsPath": v.SnapshotsPath,
		"Snapshots":     snapshots,
	})
}

func GetSnapshot(c *gin.Context) {
	volume := c.Param("volume")
	snapshot := c.Param("snapshot")
	s := models.GetSnapshot(volume, snapshot)
	err := s.LoadCacheOrFiles(s.CacheFilePath(models.GetVolumeCacheDir(volume)))
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": fmt.Sprintf("LoadCacheOrFiles error: %v", err)})
		return
	}
	c.JSON(200, s)
}

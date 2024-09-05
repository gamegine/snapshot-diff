package controllers

import (
	"snapshot-diff/models"
	"snapshot-diff/utils"

	"github.com/gin-gonic/gin"
)

var Volumes models.Volumes

func InitVolumes() error {
	volumes, err := models.LoadVolumes()
	if err != nil {
		return err
	}
	for k, v := range volumes {
		err := v.UpdateSnapshotsList()
		if err != nil {
			return err
		} else {
			CacheDirPath := v.CacheDir()
			for k, s := range v.Snapshots {
				err = s.LoadCacheOrFiles(s.CacheFilePath(CacheDirPath))
				if err != nil {
					return err
				} else {
					v.Snapshots[k] = s
				}
			}
			volumes[k] = v
		}
	}
	Volumes = volumes
	return nil
}

func GetVolumes(c *gin.Context) {
	c.JSON(200, gin.H{
		"volumes": utils.MapKeys(Volumes),
	})
}

func GetVolume(c *gin.Context) {
	volume := c.Param("volume")
	if !utils.MapContains(Volumes, volume) {
		c.JSON(404, gin.H{"code": "404", "msg": "volume not found"})
	} else {
		c.JSON(200, gin.H{
			"SnapshotsPath": Volumes[volume].SnapshotsPath,
			"Snapshots":     utils.MapKeys(Volumes[volume].Snapshots),
		})
	}
}

func GetSnapshot(c *gin.Context) {
	volume := c.Param("volume")
	snapshot := c.Param("snapshot")
	if !utils.MapContains(Volumes, volume) {
		c.JSON(404, gin.H{"code": "404", "msg": "volume not found"})
	} else {
		if !utils.MapContains(Volumes[volume].Snapshots, snapshot) {
			c.JSON(404, gin.H{"code": "404", "msg": "snapshot not found"})
		} else {
			c.JSON(200, Volumes[volume].Snapshots[snapshot])
		}
	}
}

func UpdateVolumes(c *gin.Context) {
	err := InitVolumes()
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": err})
		return
	}
	c.JSON(200, gin.H{"code": "200", "msg": "ok"})
}

func UpdateVolume(c *gin.Context) {
	volume := c.Param("volume")
	if !utils.MapContains(Volumes, volume) {
		c.JSON(404, gin.H{"code": "404", "msg": "volume not found"})
		return
	}
	v := Volumes[volume]
	err := v.UpdateSnapshotsList()
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": err})
		return
	}
	Volumes[volume] = v
	c.JSON(200, gin.H{"code": "200", "msg": "ok"})
}

func UpdateSnapshot(c *gin.Context) {
	volume := c.Param("volume")
	snapshot := c.Param("snapshot")
	if !utils.MapContains(Volumes, volume) {
		c.JSON(404, gin.H{"code": "404", "msg": "volume not found"})
		return
	}
	if !utils.MapContains(Volumes[volume].Snapshots, snapshot) {
		c.JSON(404, gin.H{"code": "404", "msg": "snapshot not found"})
		return
	}
	s := Volumes[volume].Snapshots[snapshot]
	err := s.LoadFiles()
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "msg": err})
		return
	}
	Volumes[volume].Snapshots[snapshot] = s
	c.JSON(200, gin.H{"code": "200", "msg": "ok"})
}

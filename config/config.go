package config

import (
	"os"
	"snapshot-diff/models"
)

func Config() {
	_, present := os.LookupEnv("test")
	if present {
		models.SnapshotsPath = "testdata/"
	}
}

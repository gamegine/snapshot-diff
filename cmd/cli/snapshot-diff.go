package main

import (
	"fmt"
	"snapshot-diff/config"
	"snapshot-diff/models"
	"snapshot-diff/utils"
)

func main() {
	config.Config()

	volumes, err := models.LoadVolumes()
	if err != nil {
		fmt.Printf("Error LoadVolumes: %v\n", err)
		return
	}
	fmt.Printf("volumes: %s\n", utils.MapKeys(volumes))

	for _, volume := range volumes {
		cachePath := volume.CacheDir()
		snapshots, err := volume.GetSnapshots()
		if err != nil {
			fmt.Printf("Error GetSnapshots: %v\n", err)
			return
		}
		fmt.Printf(" - volume %s, %d snapshots\n", volume.Name(), len(snapshots))
		for _, snapshotStr := range snapshots {
			snapshot := volume.GetSnapshot(snapshotStr)
			fmt.Printf("\tsnapshot: %s \n\t\tpath: %s\n\t\tcache: %s\n", snapshot.Name(), snapshot.Path, snapshot.CacheFilePath(cachePath))
			err = snapshot.LoadCacheOrFiles(snapshot.CacheFilePath(cachePath))
			if err != nil {
				fmt.Printf("Error Snapshots.LoadCacheOrFiles: %v\n", err)
				return
			}

			// Hash
			if !snapshot.IsHash() {
				err := snapshot.HashProgress()
				if err != nil {
					fmt.Printf("Error HashProgress %v\n", err)
				}
			}
			// Print
			for _, f := range snapshot.Files {
				fmt.Printf("\t\t\t%s\n", f.Path)
				if !f.IsDir && !f.IsSymlink {
					fmt.Printf("\t\t\t\tsha256:%s\n", f.Sha256)
				}
			}
			err = snapshot.SaveCache(snapshot.CacheFilePath(cachePath))
			if err != nil {
				fmt.Println("Error SaveCache ", err)
			}
		}
	}
}

package main

import (
	"fmt"
	"os"
	"snapshot-diff/models"
	"snapshot-diff/utils"
)

func main() {
	{
		_, present := os.LookupEnv("test")
		if present {
			models.SnapshotsPath = "testdata/"
		}
	}
	volumes, err := models.LoadVolumes()
	if err != nil {
		fmt.Printf("Error LoadVolumes: %v\n", err)
		return
	}
	fmt.Printf("volumes: %s\n", utils.MapKeys(volumes))

	for volumeIndex, volume := range volumes {
		CacheDirPath := volume.CacheDir()
		err := volume.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf(" - volume %s, %d snapshots\n", volume.Name(), len(volume.Snapshots))
		for snapshotIndex := range volume.Snapshots {
			snapshot := volume.Snapshots[snapshotIndex]
			fmt.Printf("\tsnapshot: %s \n\t\tpath: %s\n\t\tcache: %s\n", snapshot.Name(), snapshot.Path, snapshot.CacheFilePath(CacheDirPath))
			cacheErr := snapshot.LoadCache(snapshot.CacheFilePath(CacheDirPath))
			if cacheErr != nil {
				fmt.Println("LoadFiles")
				err = snapshot.LoadFiles()
				if err != nil {
					fmt.Printf("Error Snapshots.LoadFiles: %v\n", err)
					return
				}
				err = snapshot.LoadFilesInfo()
				if err != nil {
					fmt.Printf("Error Snapshots.LoadFilesInfo: %v\n", err)
					return
				}
				for i := range snapshot.Files {
					f := &snapshot.Files[i]
					if !f.IsDir && !f.IsSymlink {
						fmt.Println(f.Path)
						f.HashProgress()
					}
				}
			}
			for _, f := range snapshot.Files {
				fmt.Printf("\t\t\t%s\n", f.Path)
				if !f.IsDir && !f.IsSymlink {
					fmt.Printf("\t\t\t\t%s\n", f.Sha256)
				}
			}
			err = snapshot.SaveCache(snapshot.CacheFilePath(CacheDirPath))
			if err != nil {
				fmt.Println(err)
			}
			volume.Snapshots[snapshotIndex] = snapshot
		}
		volumes[volumeIndex] = volume
	}
}

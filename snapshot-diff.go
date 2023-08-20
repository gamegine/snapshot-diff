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
	/*
		for i, v := range volumes {
			v.Snapshots = append(v.Snapshots, models.Snapshot{Path: "../testdata/snapshot/GMT+01_2023-05-08_1140/"})
			volumes[i] = v
		}
	*/

	for i, v := range volumes {
		CacheDirPath := v.CacheDir()
		err := v.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf(" - volume %s, %d snapshots\n", v.Name(), len(v.Snapshots))
		for si := range v.Snapshots {
			s := &v.Snapshots[si]
			fmt.Printf("\tsnapshot: %s \n\t\tpath: %s\n\t\tcache: %s\n", s.Name(), s.Path, s.CacheFilePath(CacheDirPath))
			cacheErr := s.LoadCache(s.CacheFilePath(CacheDirPath))
			if cacheErr != nil {
				fmt.Println("load")
				err = s.LoadFiles()
				if err != nil {
					fmt.Printf("Error Snapshots.LoadFiles: %v\n", err)
					return
				}
				err = s.LoadFilesInfo()
				if err != nil {
					fmt.Printf("Error Snapshots.LoadFilesInfo: %v\n", err)
					return
				}
				for i := range s.Files {
					f := &s.Files[i]
					if !f.IsDir && !f.IsSymlink {
						fmt.Println(f.Path)
						f.HashProgress()
					}
				}
			}
			for i := range s.Files {
				f := &s.Files[i]
				fmt.Printf("\t\t\t%s\n", f.Path)
				if !f.IsDir && !f.IsSymlink {
					fmt.Printf("\t\t\t\t%s\n", f.Sha256)
				}
			}
			err = s.SaveCache(s.CacheFilePath(CacheDirPath))
			if err != nil {
				fmt.Println(err)
			}
		}
		volumes[i] = v
	}
}

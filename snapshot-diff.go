package main

import (
	"fmt"
	"os"
	"snapshot-diff/models"
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
	for i := range volumes {
		v := &volumes[i]
		CacheDirPath := v.CacheDir()
		err := v.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf("volume %s, %d snapshots\n", v.Name(), len(v.Snapshots))
		for si := range v.Snapshots {
			s := &v.Snapshots[si]
			fmt.Printf("snapshot: %s \n\tpath: %s\n\tcache: %s\n", s.Name(), s.Path, s.CacheFilePath(CacheDirPath))
			err = s.Load()
			if err != nil {
				fmt.Printf("Error Snapshots.Load: %v\n", err)
				return
			}
			err = s.LoadFilesInfo()
			if err != nil {
				fmt.Printf("Error Snapshots.LoadFilesInfo: %v\n", err)
				return
			}
			for i := range s.Files {
				f := &s.Files[i]
				fmt.Printf("\t\t%s\n", f.Path)
				if !f.IsDir {
					f.Hash()
					fmt.Printf("\t\t\t%s\n", f.Sha256)
				}

			}
			err = s.SaveCache(s.CacheFilePath(CacheDirPath))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

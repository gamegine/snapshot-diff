package main

import (
	"fmt"
	"snapshot-diff/models"
)

func main() {
	models.SnapshotsPath = "testdata/"
	volumes, err := models.LoadVolumes()
	if err != nil {
		fmt.Printf("Error LoadVolumes: %v\n", err)
		return
	}
	if len(volumes) != 0 {
		v := &volumes[0]
		err := v.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf("volume %s %d snapshots\n", v.Name(), len(v.Snapshots))
		if len(v.Snapshots) != 0 {
			s := &v.Snapshots[0]
			fmt.Printf("snapshot: %s \n\tpath: %s\n\tcache: %s\n", s.Name(), s.Path, s.CachePath(v.Name()))
			err = s.Load()
			if err != nil {
				fmt.Printf("Error Snapshots.Load: %v\n", err)
				return
			}
			for _, f := range s.Files {
				fmt.Printf("\t\t%s\n", f.Path)
			}
		}
	}
}

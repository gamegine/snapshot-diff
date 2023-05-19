package main

import (
	"fmt"
	"snapshot-diff/models"
)

func main() {
	models.SnapshotsPath = "testdata/"
	vs, err := models.LoadVolumes()
	if err != nil {
		fmt.Printf("Error LoadVolumes: %v\n", err)
		return
	}
	fmt.Println(vs)
	if len(vs) != 0 {
		v := vs[0]
		err := v.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf("volume %s %d snapshots\n", v.Name(), len(v.Snapshots))
		if len(v.Snapshots) != 0 {
			fmt.Printf("snapshot 1 path: %s, cache: %s\n", v.Snapshots[0].Path, v.Snapshots[0].CachePath(v.Name()))
			err = v.Snapshots[0].Load()
			if err != nil {
				fmt.Printf("Error Snapshots.Load: %v\n", err)
				return
			}
			for _, f := range v.Snapshots[0].Files {
				fmt.Printf("\t%s\n", f.Path)
			}
		}
	}
}

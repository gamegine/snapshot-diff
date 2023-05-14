package main

import (
	"fmt"
	"snapshot-diff/models"
)

func printSnapshot(s models.Snapshot) {
	fmt.Printf("path:%s\nfiles:\n", s.Path)
	for _, f := range s.Files {
		fmt.Printf("\t%s\n", f.Path)
	}
}

func main() {
	models.SnapshotsPath = "./testdata"
	vs, err := models.LoadVolumes()
	if err != nil {
		fmt.Printf("Error LoadVolumes: %v\n", err)
		return
	}
	fmt.Print(vs)
	if len(vs) != 0 {
		v := vs[0]
		err := v.UpdateSnapshotsList()
		if err != nil {
			fmt.Printf("Error UpdateSnapshotsList: %v\n", err)
			return
		}
		fmt.Printf(" %d snapshots\n", len(v.Snapshots))
		if len(v.Snapshots) != 0 {
			fmt.Printf("snapshot 1 path:%s\n", v.Snapshots[0].Path)
			err = v.Snapshots[0].Load()
			if err != nil {
				fmt.Printf("Error Snapshots.Load: %v\n", err)
				return
			}
			printSnapshot(v.Snapshots[0])
		}
	}
}

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

	v, _ := models.LoadVolumes()
	fmt.Println(v)
	_ = v[0].UpdateSnapshotsList()
	printSnapshot(v[0].Snapshots[0])
	_ = v[0].Snapshots[0].Load()
	printSnapshot(v[0].Snapshots[0])
}

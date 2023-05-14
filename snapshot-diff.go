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
	var s = models.Snapshot{Path: "./testdata/snapshot/GMT+01_2023-05-08_1140/"}
	s.Load()
	s.LoadFiles()
	printSnapshot(s)
}

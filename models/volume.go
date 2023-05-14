package models

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Volume struct {
	SnapshotsPath string
	Snapshots     Snapshots
}

type Volumes []Volume

func RemoveFiles(slice []fs.DirEntry) []fs.DirEntry {
	var sliceres []fs.DirEntry
	for _, entry := range slice {
		if entry.IsDir() {
			sliceres = append(sliceres, entry)
		}
	}
	return sliceres
}

func (v *Volume) UpdateSnapshotsList() error {
	entries, err := os.ReadDir(v.SnapshotsPath)
	if err != nil {
		return err
	}
	entries = RemoveFiles(entries)
	for _, e := range entries {
		v.Snapshots = append(v.Snapshots, Snapshot{Path: filepath.Join(v.SnapshotsPath, e.Name())})
	}
	return nil
}

var SnapshotsPath = "/mnt/snapshot/export/Unified-Snapshot/"

func LoadVolumes() (Volumes, error) {
	var v Volumes
	entries, err := os.ReadDir(SnapshotsPath)
	if err != nil {
		return nil, err
	}
	entries = RemoveFiles(entries)
	for _, e := range entries {
		v = append(v, Volume{SnapshotsPath: filepath.Join(SnapshotsPath, e.Name())})
	}
	return v, nil
}

package models

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Volume struct {
	SnapshotsPath string
	Snapshots     Snapshots
}

type Volumes []Volume

func contain(v Volume, snapshotsPath string) bool {
	for _, s := range v.Snapshots {
		if s.Path == snapshotsPath {
			return true
		}
	}
	return false
}

func (v *Volume) UpdateSnapshotsList() error {
	entries, err := os.ReadDir(v.SnapshotsPath)
	if err != nil {
		return err
	}
	// entries = RemoveFiles(entries) // symlink
	for _, e := range entries {
		if !contain(*v, filepath.Join(v.SnapshotsPath, e.Name())) {
			v.Snapshots = append(v.Snapshots, Snapshot{Path: filepath.Join(v.SnapshotsPath, e.Name())})
		}
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
	for _, e := range entries {
		if e.IsDir() {
			v = append(v, Volume{SnapshotsPath: filepath.Join(SnapshotsPath, e.Name())})
		}
	}
	return v, nil
}

func (v *Volume) Name() string {
	return strings.ReplaceAll(strings.ReplaceAll(v.SnapshotsPath, SnapshotsPath, ""), " Snapshot", "")
}

var SnapshotsCachePath = "./cache"

func (v *Volume) CacheDir() string {
	path := path.Join(SnapshotsCachePath, v.Name())
	os.MkdirAll(path, os.ModePerm)
	return path
}

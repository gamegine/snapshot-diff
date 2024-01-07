package models

import (
	"os"
	"path"
	"path/filepath"
	"snapshot-diff/utils"
	"strings"
)

type Volume struct {
	SnapshotsPath string
	Snapshots     Snapshots
}

type Volumes map[string]Volume

func (v *Volume) UpdateSnapshotsList() error {
	entries, err := os.ReadDir(v.SnapshotsPath)
	if err != nil {
		return err
	}
	// Check if a map is initialised
	if v.Snapshots == nil {
		v.Snapshots = make(Snapshots)
	}
	// entries = RemoveFiles(entries) // symlink
	for _, e := range entries {
		if !utils.MapContains(v.Snapshots, e.Name()) {
			v.Snapshots[e.Name()] = Snapshot{Path: filepath.Join(v.SnapshotsPath, e.Name())}
		}
	}
	return nil
}

var SnapshotsPath = "/mnt/snapshot/export/Unified-Snapshot/"

func LoadVolumes() (Volumes, error) {
	v := make(Volumes)
	entries, err := os.ReadDir(SnapshotsPath)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() {
			v[e.Name()] = Volume{SnapshotsPath: filepath.Join(SnapshotsPath, e.Name())}
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

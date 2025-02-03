package models

import (
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

type Volume struct {
	SnapshotsPath string
}

type Volumes map[string]Volume

var SnapshotsPath = "/mnt/snapshot/export/Unified-Snapshot/"

func (v *Volume) Name() string {
	return strings.ReplaceAll(strings.ReplaceAll(v.SnapshotsPath, SnapshotsPath, ""), " Snapshot", "")
}

var SnapshotsCachePath = "./cache"

func (v *Volume) CacheDir() string {
	path := path.Join(SnapshotsCachePath, v.Name())
	os.MkdirAll(path, os.ModePerm)
	return path
}

func (v *Volume) GetSnapshots() ([]string, error) {
	entries, err := os.ReadDir(v.SnapshotsPath)
	if err != nil {
		return nil, err
	}
	var s []string
	for _, e := range entries {
		s = append(s, e.Name())
	}
	slices.Sort(s)
	return s, nil
}

func (v *Volume) GetSnapshot(snapshot string) Snapshot {
	return Snapshot{Path: filepath.Join(v.SnapshotsPath, snapshot)}
}

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

func GetVolume(volume string) Volume {
	return Volume{SnapshotsPath: filepath.Join(SnapshotsPath, volume)}
}

func GetSnapshot(volume string, snapshots string) Snapshot {
	v := GetVolume(volume)
	return v.GetSnapshot(snapshots)
}

func GetVolumeCacheDir(volume string) string {
	v := GetVolume(volume)
	return v.CacheDir()
}

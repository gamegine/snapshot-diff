package models

import (
	"os"
	"reflect"
	"snapshot-diff/utils"
	"testing"
)

func TestLoadVolumes(t *testing.T) {
	SnapshotsPath = "../testdata"
	got, err := LoadVolumes()
	if err != nil {
		t.Errorf("error %v", err)
	}

	want := Volumes{
		"volume": {SnapshotsPath: "../testdata/volume"},
	}

	if !reflect.DeepEqual(utils.MapKeys(got), utils.MapKeys(want)) {
		t.Errorf("got %v, wanted %v", utils.MapKeys(got), utils.MapKeys(want))
	}

	for k, v := range got {
		if v.SnapshotsPath != want[k].SnapshotsPath {
			t.Errorf("got %v wanted %v", v.SnapshotsPath, want[k].SnapshotsPath)
		}
	}
}

func TestLoadVolumesWithErr(t *testing.T) {
	SnapshotsPath = "./undef"
	_, err := LoadVolumes()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestVolumeContain(t *testing.T) {
	var vol = Volume{
		Snapshots: Snapshots{
			Snapshot{Path: "a"},
			Snapshot{Path: "c"},
		},
	}
	if !contain(vol, "a") {
		t.Errorf("a not in vol[a,c], wanted yes")
	}
	if contain(vol, "b") {
		t.Errorf("b in vol[a,c], wanted yes")
	}
	if !contain(vol, "c") {
		t.Errorf("c not in vol[a,c], wanted yes")
	}
}

func TestVolumeUpdateSnapshotsList(t *testing.T) {
	var got = Volume{SnapshotsPath: "../testdata/volume"}
	err := got.UpdateSnapshotsList()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"../testdata/volume/snapshot",
		"../testdata/volume/symlink",
	}

	if len(got.Snapshots) != len(want) {
		t.Errorf("got %v, wanted %v", len(got.Snapshots), len(want))
		return
	}
	for i, s := range got.Snapshots {
		if s.Path != want[i] {
			t.Errorf("got %v wanted %v", s.Path, want[i])
		}
	}
}

func TestVolumeUpdateSnapshotsListWithExistingSnapshot(t *testing.T) {
	var got = Volume{
		SnapshotsPath: "../testdata/volume",
		Snapshots: Snapshots{
			Snapshot{
				Path:  "../testdata/volume/snapshot",
				Files: Files{File{}},
			},
		},
	}
	err := got.UpdateSnapshotsList()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"../testdata/volume/snapshot",
		"../testdata/volume/symlink",
	}

	if len(got.Snapshots) != len(want) {
		t.Errorf("got %v, wanted %v", len(got.Snapshots), len(want))
		return
	}
	for i, s := range got.Snapshots {
		if s.Path != want[i] {
			t.Errorf("got %v wanted %v", s.Path, want[i])
		}
		if s.Path == "../testdata/volume/snapshot" {
			if len(s.Files) != 1 {
				t.Errorf("existing snapshot data was erased %v", s)
			}
		}
	}
}

func TestVolumeUpdateSnapshotsListWithErr(t *testing.T) {
	var v = Volume{SnapshotsPath: "./undef"}
	err := v.UpdateSnapshotsList()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestVolumeName(t *testing.T) {
	SnapshotsPath = "../testdata/"
	var s = Volume{SnapshotsPath: "../testdata/snapshot Snapshot"}
	got := s.Name()
	want := "snapshot"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCacheDir(t *testing.T) {
	SnapshotsPath = "../testdata/"
	SnapshotsCachePath = "../cache"
	var s = Volume{SnapshotsPath: "../testdata/snapshot Snapshot"}
	got := s.CacheDir()
	want := "../cache/snapshot"
	if got != want {
		t.Errorf("got cache path %v, wanted %v", got, want)
	}
	if _, err := os.Stat(got); os.IsNotExist(err) {
		t.Error("cache path does not exist")
	}
}

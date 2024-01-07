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

func TestVolumeUpdateSnapshotsList(t *testing.T) {
	var got = Volume{SnapshotsPath: "../testdata/volume"}
	err := got.UpdateSnapshotsList()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"snapshot",
		"symlink",
	}
	if len(got.Snapshots) != len(want) {
		t.Errorf("got %v, wanted %v", len(got.Snapshots), len(want))
		return
	}
	for _, s := range want {
		if !utils.MapContains(got.Snapshots, s) {
			t.Errorf("got %v wanted %v", utils.MapKeys(got.Snapshots), s)
		}
		if got.Snapshots[s].Path != "../testdata/volume/"+s {
			t.Errorf("got %v wanted ../testdata/volume/%v", got.Snapshots[s].Path, s)
		}
	}
}

func TestVolumeUpdateSnapshotsListWithExistingSnapshot(t *testing.T) {
	var got = Volume{
		SnapshotsPath: "../testdata/volume",
		Snapshots: Snapshots{
			"snapshot": Snapshot{
				Path:  "../testdata/volume/snapshot",
				Files: Files{File{Path: "test"}},
			},
		},
	}
	err := got.UpdateSnapshotsList()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"snapshot",
		"symlink",
	}
	if len(got.Snapshots) != len(want) {
		t.Errorf("got %v, wanted %v", len(got.Snapshots), len(want))
		return
	}
	for _, s := range want {
		if !utils.MapContains(got.Snapshots, s) {
			t.Errorf("got %v wanted %v", utils.MapKeys(got.Snapshots), s)
		}
		if got.Snapshots[s].Path != "../testdata/volume/"+s {
			t.Errorf("got %v wanted ../testdata/volume/%v", got.Snapshots[s].Path, s)
		}
	}
	// snapshot data not deleted
	if len(got.Snapshots["snapshot"].Files) != 1 {
		t.Errorf("got %v wanted %v", len(got.Snapshots["snapshot"].Files), 1)
	}
	if got.Snapshots["snapshot"].Files[0].Path != "test" {
		t.Errorf("got %v wanted %v", got.Snapshots["snapshot"].Files[0].Path, "test")
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

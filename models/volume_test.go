package models

import (
	"os"
	"reflect"
	"snapshot-diff/utils"
	"testing"
)

func TestLoadVolumes(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name     string
		path     string
		error    bool
		Expected Volumes
	}{
		{
			name:     "volume",
			path:     "../testdata",
			error:    false,
			Expected: Volumes{"volume": {SnapshotsPath: "../testdata/volume"}},
		},
		{
			name:     "path error",
			path:     "./undef",
			error:    true,
			Expected: Volumes{},
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			SnapshotsPath = TestCase.path
			got, err := LoadVolumes()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
			if !reflect.DeepEqual(utils.MapKeys(got), utils.MapKeys(TestCase.Expected)) {
				t.Errorf("got %v, wanted %v", utils.MapKeys(got), utils.MapKeys(TestCase.Expected))
			}
			for k, v := range got {
				if v.SnapshotsPath != TestCase.Expected[k].SnapshotsPath {
					t.Errorf("got %v wanted %v", v.SnapshotsPath, TestCase.Expected[k].SnapshotsPath)
				}
			}
		})
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

func TestVolumeGetVolume(t *testing.T) {
	SnapshotsPath = "../testdata/"
	got := GetVolume("volume")
	want := Volume{SnapshotsPath: "../testdata/volume"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestVolumeGetSnapshots(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name     string
		volume   Volume
		error    bool
		Expected []string
	}{
		{
			name:     "volume",
			volume:   Volume{SnapshotsPath: "../testdata/volume"},
			error:    false,
			Expected: []string{"snapshot", "symlink"},
		},
		{
			name:     "error",
			volume:   Volume{SnapshotsPath: "./undef"},
			error:    true,
			Expected: nil,
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			snapshots, err := TestCase.volume.GetSnapshots()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
			if !reflect.DeepEqual(snapshots, TestCase.Expected) {
				t.Errorf("got %v, wanted %v", snapshots, TestCase.Expected)
			}
		})
	}
}

func TestVolumeGetSnapshot(t *testing.T) {
	SnapshotsPath = "../testdata/"
	volume := GetVolume("volume")

	got := volume.GetSnapshot("snapshot")
	want := Snapshot{Path: "../testdata/volume/snapshot"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestVolumeGetSnapshot1(t *testing.T) {
	SnapshotsPath = "../testdata/"

	got := GetSnapshot("volume", "snapshot")
	want := Snapshot{Path: "../testdata/volume/snapshot"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestVolumeGetVolumeCacheDir(t *testing.T) {
	SnapshotsPath = "../testdata/"
	SnapshotsCachePath = "../cache"

	got := GetVolumeCacheDir("volume")
	want := "../cache/volume"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

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

func TestVolumeUpdateSnapshotsList(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name          string
		volume        Volume
		error         bool
		Expected      []string
		ExpectedFiles []string
	}{
		{
			name:          "volume",
			volume:        Volume{SnapshotsPath: "../testdata/volume"},
			error:         false,
			Expected:      []string{"snapshot", "symlink"},
			ExpectedFiles: []string{},
		},
		{
			name: "existing snapshot",
			volume: Volume{SnapshotsPath: "../testdata/volume",
				Snapshots: Snapshots{
					"snapshot": Snapshot{
						Path:  "../testdata/volume/snapshot",
						Files: Files{File{Path: "test"}},
					},
				},
			},
			error:         false,
			Expected:      []string{"snapshot", "symlink"},
			ExpectedFiles: []string{"test"},
		},
		{
			name:          "error",
			volume:        Volume{SnapshotsPath: "./undef"},
			error:         true,
			Expected:      []string{},
			ExpectedFiles: []string{},
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			err := TestCase.volume.UpdateSnapshotsList()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
			if len(TestCase.volume.Snapshots) != len(TestCase.Expected) {
				t.Errorf("got %v, wanted %v", len(TestCase.volume.Snapshots), len(TestCase.Expected))
				return
			}
			for _, s := range TestCase.Expected {
				if !utils.MapContains(TestCase.volume.Snapshots, s) {
					t.Errorf("got %v wanted %v", utils.MapKeys(TestCase.volume.Snapshots), s)
				}
				if TestCase.volume.Snapshots[s].Path != TestCase.volume.SnapshotsPath+"/"+s {
					t.Errorf("got %v wanted %s/%v", TestCase.volume.Snapshots[s].Path, TestCase.volume.SnapshotsPath, s)
				}
			}

			// snapshot data not deleted
			if len(TestCase.volume.Snapshots["snapshot"].Files) != len(TestCase.ExpectedFiles) {
				t.Errorf("got %v wanted %v", len(TestCase.volume.Snapshots["snapshot"].Files), 1)
			}
			for i := range TestCase.ExpectedFiles {
				if TestCase.volume.Snapshots["snapshot"].Files[i].Path != TestCase.ExpectedFiles[i] {
					t.Errorf("got %v wanted %v", TestCase.volume.Snapshots["snapshot"].Files[i].Path, TestCase.ExpectedFiles[i])
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

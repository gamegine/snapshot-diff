package models

import (
	"testing"
)

func TestLoadVolumes(t *testing.T) {
	SnapshotsPath = "../testdata"
	got, err := LoadVolumes()
	if err != nil {
		t.Errorf("error %v", err)
	}

	want := Volumes{{SnapshotsPath: "../testdata/snapshot"}}

	// if reflect.DeepEqual(want, got) {
	// 	t.Errorf("got %v, wanted %v", got, want)
	// }
	if len(got) != len(want) {
		t.Errorf("got %v, wanted %v", len(got), len(want))
	}
	for i, v := range got {
		if v.SnapshotsPath != want[i].SnapshotsPath {
			t.Errorf("got %v wanted %v", v.SnapshotsPath, want[i].SnapshotsPath)
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
	var got = Volume{SnapshotsPath: "../testdata/snapshot"}
	err := got.UpdateSnapshotsList()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"../testdata/snapshot/GMT+01_2023-05-08_1034",
		"../testdata/snapshot/GMT+01_2023-05-08_1140",
		"../testdata/snapshot/GMT+01_2023-05-08_1142",
		"../testdata/snapshot/GMT+01_2023-05-08_1144",
		"../testdata/snapshot/GMT+01_2023-05-08_1145",
	}

	if len(got.Snapshots) != len(want) {
		t.Errorf("got %v, wanted %v", len(got.Snapshots), len(want))
	}
	for i, s := range got.Snapshots {
		if s.Path != want[i] {
			t.Errorf("got %v wanted %v", s.Path, want[i])
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

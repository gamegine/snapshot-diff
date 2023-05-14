package models

import (
	"testing"
)

func TestSnapshotLoad(t *testing.T) {
	var got = Snapshot{Path: "../testdata/snapshot/GMT+01_2023-05-08_1140/"}
	err := got.Load()
	if err != nil {
		t.Errorf("error %v", err)
	}
	paths := []string{
		"../testdata/snapshot/GMT+01_2023-05-08_1140/.qextension",
		"../testdata/snapshot/GMT+01_2023-05-08_1140/snapshots",
		"../testdata/snapshot/GMT+01_2023-05-08_1140/snapshots/@Recycle",
		"../testdata/snapshot/GMT+01_2023-05-08_1140/snapshots/@Recycle/desktop.ini",
		"../testdata/snapshot/GMT+01_2023-05-08_1140/snapshots/test.txt",
	}

	if len(paths) != len(got.Files) {
		t.Errorf("got %v, wanted %v", len(paths), len(got.Files))
	}

	// []gotpaths = got.[]Files.Path
	gotpaths := make([]string, len(got.Files))
	for i, f := range got.Files {
		gotpaths[i] = f.Path
	}

	// if reflect.DeepEqual(paths, gotpaths) {
	// 	t.Errorf("got \n%v\nwanted \n%v", paths, gotpaths)
	// }
	for i, p := range gotpaths {
		if p != paths[i] {
			t.Errorf("got %v wanted %v", p, paths[i])
		}
	}
}

func TestSnapshotLoadWithErr(t *testing.T) {
	var s = Snapshot{Path: "undef"}
	err := s.Load()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestSnapshotLoadFiles(t *testing.T) {
	s := Snapshot{Files: []File{{Path: "./file.go"}}}
	err := s.LoadFiles()
	if err != nil {
		t.Errorf("LoadFiles error %v", err)
	}
	if s.Files[0].Chmod == 0 {
		t.Errorf("file is not load %v", s.Files[0])
	}
}

func TestSnapshotLoadFilesWithErr(t *testing.T) {
	s := Snapshot{Files: []File{
		{Path: "./file.go"},
		{Path: "./undef"},
	}}
	err := s.LoadFiles()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

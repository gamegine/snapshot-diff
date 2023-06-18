package models

import (
	"testing"
	"time"
)

func TestSnapshotLoadFiles(t *testing.T) {
	var got = Snapshot{Path: "../testdata/snapshot/GMT+01_2023-05-08_1140/"}
	err := got.LoadFiles()
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
		return
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

func TestSnapshotLoadFilesAlreadyloaded(t *testing.T) {
	var got = Snapshot{
		Path:  "../testdata/snapshot/GMT+01_2023-05-08_1140/",
		Files: Files{File{Path: "file"}},
	}
	err := got.LoadFiles()
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(got.Files) != 1 {
		t.Errorf("loaded files not preserved %v", got.Files)
		return
	}
	if got.Files[0].Path != "file" {
		t.Errorf("loaded files not preserved %v", got.Files[0])
		return
	}
}

func TestSnapshotLoadFilesSymlink(t *testing.T) {
	var got = Snapshot{Path: "../testdata/snapshot/symlink"}
	err := got.LoadFiles()
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := []string{
		"../testdata/snapshot/GMT+01_2023-05-08_1034/.qextension",
		"../testdata/snapshot/GMT+01_2023-05-08_1034/snapshots",
		"../testdata/snapshot/GMT+01_2023-05-08_1034/snapshots/@Recycle",
		"../testdata/snapshot/GMT+01_2023-05-08_1034/snapshots/@Recycle/desktop.ini",
	}

	if len(want) != len(got.Files) {
		t.Errorf("got %v, wanted %v", len(want), len(got.Files))
		return
	}

	// []gotpaths = got.[]Files.Path
	gotpaths := make([]string, len(got.Files))
	for i, f := range got.Files {
		gotpaths[i] = f.Path
	}
	for i, p := range gotpaths {
		if p != want[i] {
			t.Errorf("got %v wanted %v", p, want[i])
		}
	}
}

func TestSnapshotLoadFilesWithErr(t *testing.T) {
	var s = Snapshot{Path: "undef"}
	err := s.LoadFiles()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestSnapshotLoadFilesInfo(t *testing.T) {
	s := Snapshot{Files: []File{{Path: "./file.go"}}}
	err := s.LoadFilesInfo()
	if err != nil {
		t.Errorf("LoadFilesInfo error %v", err)
	}
	if s.Files[0].Mode == 0 {
		t.Errorf("file is not load %v", s.Files[0])
	}
}

func TestSnapshotLoadFilesInfoWithErr(t *testing.T) {
	s := Snapshot{Files: []File{
		{Path: "./file.go"},
		{Path: "./undef"},
	}}
	err := s.LoadFilesInfo()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestSnapshotCacheFilePath(t *testing.T) {
	s := Snapshot{Path: "testdata/snapshot/GMT+01_2023-05-08_1034"}
	got := s.CacheFilePath("cache")
	want := "cache/GMT+01_2023-05-08_1034.json"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSnapshotSaveCache(t *testing.T) {
	s := Snapshot{Files: []File{
		{Path: "./file.go"},
	}}
	err := s.SaveCache("../testdata/save-snapshot.json")
	if err != nil {
		t.Errorf("SaveCache error %v", err)
	}
}

func TestSnapshotSaveCacheWithPathErr(t *testing.T) {
	s := Snapshot{Files: []File{
		{Path: "./file.go"},
	}}
	err := s.SaveCache("./")
	if err == nil {
		t.Error("no error with dir path to save file")
	}
}

func TestSnapshotLoadCache(t *testing.T) {
	s := Snapshot{}
	err := s.LoadCache("../testdata/snapshot-cache.json")
	if err != nil {
		t.Errorf("LoadCache error %v", err)
	}

	T, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	want := Snapshot{
		Path: "testdata/snapshot/GMT+01_2023-05-08_1034",
		Files: Files{
			{Path: "testdata/snapshot/GMT+01_2023-05-08_1034/.qextension", IsDir: false, Mode: 0, Size: 0, ModifTime: T},
			{Path: "testdata/snapshot/GMT+01_2023-05-08_1034/snapshots", IsDir: false, Mode: 0, Size: 0, ModifTime: T},
			{Path: "testdata/snapshot/GMT+01_2023-05-08_1034/snapshots/@Recycle", IsDir: false, Mode: 0, Size: 0, ModifTime: T},
			{Path: "testdata/snapshot/GMT+01_2023-05-08_1034/snapshots/@Recycle/desktop.ini", IsDir: false, Mode: 0, Size: 0, ModifTime: T},
		}}
	_ = want
	if s.Path != want.Path {
		t.Errorf("got %v, wanted %v", s, want)
	}
	if len(s.Files) != len(want.Files) {
		t.Errorf("got %v, wanted %v", s, want)
	}
	for i := range s.Files {
		if s.Files[i] != want.Files[i] {
			t.Errorf("got %v, wanted %v", s, want)
		}
	}

}

func TestSnapshotLoadCacheWithPathErr(t *testing.T) {
	s := Snapshot{}
	err := s.LoadCache("./undef")
	if err == nil {
		t.Error("no error with dir path to save file")
	}
}

func TestSnapshotLoadCacheWithJsonErr(t *testing.T) {
	s := Snapshot{}
	err := s.LoadCache("../testdata/snapshot-cache-error.json")
	if err == nil {
		t.Error("no error with invalid json format")
	}
}

func TestSnapshotName(t *testing.T) {
	s := Snapshot{Path: "testdata/snapshot/GMT+01_2023-05-08_1034"}
	got := s.Name()
	want := "GMT+01_2023-05-08_1034"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

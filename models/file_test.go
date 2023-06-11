package models

import (
	"os"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	var got = File{Path: "./file.go"}
	err := got.Load()
	if err != nil {
		t.Errorf("error %v", err)
	}

	fi, err := os.Stat("./file.go")
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := File{Path: "./file.go", IsDir: false, Chmod: uint32(fi.Mode()), Size: fi.Size(), ModifTime: fi.ModTime()}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNewFileWithErr(t *testing.T) {
	var f = File{Path: "./undef"}
	err := f.Load()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestHash(t *testing.T) {
	var f = File{Path: "../LICENSE"}
	err := f.Hash()
	if err != nil {
		t.Errorf("error %v", err)
	}
	got := f.Sha256
	want := "6634449D791CDB054AD21E4602AB0E0912DC3B1629DF90CB08512347D35F53E9"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestHashWithErr(t *testing.T) {
	var f = File{Path: "./undef"}
	err := f.Hash()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestHashProgress(t *testing.T) {
	var f = File{Path: "../LICENSE"}
	err := f.HashProgress()
	if err != nil {
		t.Errorf("error %v", err)
	}
	got := f.Sha256
	want := "6634449D791CDB054AD21E4602AB0E0912DC3B1629DF90CB08512347D35F53E9"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestHashProgressWithErr(t *testing.T) {
	var f = File{Path: "./undef"}
	err := f.HashProgress()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestDiff(t *testing.T) {
	a := File{Path: "./file.txt",
		IsDir:     false,
		Sha256:    "sha",
		Chmod:     12,
		Size:      32,
		ModifTime: time.Now(),
	}
	b := File{Path: "./file.txt",
		IsDir:     false,
		Sha256:    "sha",
		Chmod:     12,
		Size:      32,
		ModifTime: time.Now(),
	}

	if Diff(a, b) {
		t.Errorf("files is not diff")
	}
	//
	b.IsDir = true
	if !Diff(a, b) {
		t.Errorf("files is diff")
	}
	b.IsDir = false
	//
	b.Size = 31
	if !Diff(a, b) {
		t.Errorf("files is diff")
	}
	b.Size = 32
	//
	b.Sha256 = "---"
	if !Diff(a, b) {
		t.Errorf("files is diff")
	}
	b.Sha256 = "sha"
}

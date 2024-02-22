package models

import (
	"io/fs"
	"os"
	"testing"
	"time"
)

func TestLoadFileInfo(t *testing.T) {
	fi, err := os.Stat("./file.go")
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := File{Path: "./file.go", IsDir: false, Mode: fi.Mode(), Size: fi.Size(), ModifTime: fi.ModTime()}

	var got = File{Path: "./file.go"}
	got.LoadFileInfo(fi)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoadFile(t *testing.T) {
	var got = File{Path: "./file.go"}
	err := got.Load()
	if err != nil {
		t.Errorf("error %v", err)
	}

	fi, err := os.Stat("./file.go")
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := File{Path: "./file.go", IsDir: false, Mode: fi.Mode(), Size: fi.Size(), ModifTime: fi.ModTime()}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoadFileWithErr(t *testing.T) {
	var f = File{Path: "./undef"}
	err := f.Load()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestIsSpecialFile(t *testing.T) {
	var f = File{Path: "../LICENSE"}
	err := f.Load()
	if err != nil {
		t.Errorf("error %v", err)
	}
	if IsSpecialFile(f) {
		t.Errorf("LICENSE is not special file")
	}
	f.IsDir = true
	if !IsSpecialFile(f) {
		t.Errorf("is folder")
	}
	f.IsDir = false

	m := f.Mode
	f.Mode |= fs.ModeSocket
	if !IsSpecialFile(f) {
		t.Errorf("Socket is special file")
	}
	f.Mode = m
	f.Mode |= fs.ModeNamedPipe
	if !IsSpecialFile(f) {
		t.Errorf("Pipe is special file")
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

func TestHashWithSpecialFile(t *testing.T) {
	var f = File{Path: "../LICENSE", Mode: fs.ModeSocket}
	err := f.Hash()
	if err != nil {
		t.Errorf("error %v", err)
	}
	got := f.Sha256
	want := ""
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

func TestHashProgressWithSpecialFile(t *testing.T) {
	var f = File{Path: "../LICENSE", Mode: fs.ModeSocket}
	err := f.HashProgress()
	if err != nil {
		t.Errorf("error %v", err)
	}
	got := f.Sha256
	want := ""
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
		Mode:      12,
		Size:      32,
		ModifTime: time.Now(),
	}
	b := File{Path: "./file.txt",
		IsDir:     false,
		Sha256:    "sha",
		Mode:      12,
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

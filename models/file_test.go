package models

import (
	"os"
	"testing"
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

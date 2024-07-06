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
	want := File{APath: "./file.go", Path: "./file.go", IsDir: false, Mode: fi.Mode(), Size: fi.Size(), ModifTime: fi.ModTime()}

	var got = File{APath: "./file.go", Path: "./file.go"}
	got.LoadFileInfo(fi)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoadFile(t *testing.T) {
	var got = File{APath: "./file.go", Path: "./file.go"}
	err := got.Load()
	if err != nil {
		t.Errorf("error %v", err)
	}

	fi, err := os.Stat("./file.go")
	if err != nil {
		t.Errorf("error %v", err)
	}
	want := File{APath: "./file.go", Path: "./file.go", IsDir: false, Mode: fi.Mode(), Size: fi.Size(), ModifTime: fi.ModTime()}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoadFileWithErr(t *testing.T) {
	var f = File{APath: "./undef", Path: "./undef"}
	err := f.Load()
	if err == nil {
		t.Errorf("no error with undefined file")
	}
}

func TestIsSpecialFile(t *testing.T) {
	var f = File{APath: "../LICENSE"}
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
	// Anonymous struct of test cases
	tests := []struct {
		name     string
		file     File
		error    bool
		Expected string
	}{
		{
			name:     "hash",
			file:     File{APath: "../LICENSE"},
			error:    false,
			Expected: "6634449D791CDB054AD21E4602AB0E0912DC3B1629DF90CB08512347D35F53E9",
		},
		{
			name:     "specialFile",
			file:     File{APath: "../LICENSE", Mode: fs.ModeSocket},
			error:    false,
			Expected: "",
		},
		{
			name:     "error",
			file:     File{APath: "./undef"},
			error:    true,
			Expected: "",
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			var f = TestCase.file
			err := f.Hash()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
			got := f.Sha256
			if got != TestCase.Expected {
				t.Errorf("got %v, wanted %v", got, TestCase.Expected)
			}
		})
	}
}

func TestHashProgress(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name     string
		file     File
		error    bool
		Expected string
	}{
		{
			name:     "hash",
			file:     File{APath: "../LICENSE"},
			error:    false,
			Expected: "6634449D791CDB054AD21E4602AB0E0912DC3B1629DF90CB08512347D35F53E9",
		},
		{
			name:     "specialFile",
			file:     File{APath: "../LICENSE", Mode: fs.ModeSocket},
			error:    false,
			Expected: "",
		},
		{
			name:     "error",
			file:     File{APath: "./undef"},
			error:    true,
			Expected: "",
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			var f = TestCase.file
			err := f.HashProgress()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
			got := f.Sha256
			if got != TestCase.Expected {
				t.Errorf("got %v, wanted %v", got, TestCase.Expected)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	a := File{APath: "./file.txt",
		IsDir:     false,
		Sha256:    "sha",
		Mode:      12,
		Size:      32,
		ModifTime: time.Now(),
	}
	b := File{APath: "./file.txt",
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

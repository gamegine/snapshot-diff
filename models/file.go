package models

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

type File struct {
	Path      string      `json:"path"`
	IsDir     bool        `json:"dir"`
	IsSymlink bool        `json:"symlink"`
	Sha256    string      `json:"sha256,omitempty"`
	Mode      fs.FileMode `json:"mode"`
	Size      int64       `json:"bytes,omitempty"` // length in bytes for regular files
	ModifTime time.Time   `json:"modified"`        // modification time
}

type Files []File

func (f *File) LoadFileInfo(fileInfo os.FileInfo) {
	*f = File{
		Path:      f.Path,
		IsDir:     fileInfo.IsDir(),
		IsSymlink: fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink,
		Mode:      fileInfo.Mode(),
		Size:      fileInfo.Size(),
		ModifTime: fileInfo.ModTime(),
	}
}

func (f *File) Load() error {
	fileInfo, err := os.Lstat(f.Path)
	if err != nil {
		return err

	}
	f.LoadFileInfo(fileInfo)
	return nil
}

func IsSpecialFile(f File) bool {
	return f.IsDir || // IsSpecialFile is used avoid hash, we cannot hash a folder
		f.Mode&fs.ModeSocket == fs.ModeSocket ||
		f.Mode&fs.ModeNamedPipe == fs.ModeNamedPipe
}

func (f *File) Hash() error {
	if IsSpecialFile(*f) { // skip special files
		return nil
	}
	fo, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fo.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, fo)
	if err != nil {
		return err
	}
	f.Sha256 = fmt.Sprintf("%X", hash.Sum(nil))
	return nil
}

func (f *File) HashProgress() error {
	if IsSpecialFile(*f) { // skip special files
		return nil
	}
	fo, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fo.Close()

	bar := progressbar.DefaultBytes(
		f.Size,
		f.Path,
	)
	defer bar.Close()

	hash := sha256.New()
	_, err = io.Copy(io.MultiWriter(hash, bar), fo)
	if err != nil {
		return err
	}
	f.Sha256 = fmt.Sprintf("%X", hash.Sum(nil))
	return nil
}

func Diff(a, b File) bool {
	return a.Mode != b.Mode ||
		a.IsDir != b.IsDir ||
		a.Size != b.Size ||
		a.Sha256 != b.Sha256
}

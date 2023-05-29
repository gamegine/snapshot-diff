package models

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"time"
)

type File struct {
	Path      string    `json:"path"`
	IsDir     bool      `json:"dir"`
	Sha256    string    `json:"sha256,omitempty"`
	Chmod     uint32    `json:"chmod"`           // chmod bits
	Size      int64     `json:"bytes,omitempty"` // length in bytes for regular files
	ModifTime time.Time `json:"modified"`        // modification time
}

type Files []File

func (f *File) Load() error {
	fileInfo, err := os.Stat(f.Path)
	if err != nil {
		return err

	}
	*f = File{
		Path:      f.Path,
		IsDir:     fileInfo.IsDir(),
		Chmod:     uint32(fileInfo.Mode()),
		Size:      fileInfo.Size(),
		ModifTime: fileInfo.ModTime(),
	}
	return nil
}

func (f *File) Hash() error {
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

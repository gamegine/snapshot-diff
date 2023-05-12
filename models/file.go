package models

import "time"

type File struct {
	Path      string    `json:"path"`
	IsDir     bool      `json:"dir"`
	Sha256    string    `json:"sha256,omitempty"`
	Chmod     uint32    `json:"chmod"`           // chmod bits
	Size      int64     `json:"bytes,omitempty"` // length in bytes for regular files
	ModifTime time.Time `json:"modified"`        // modification time
}

type Files []File

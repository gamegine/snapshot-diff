package models

type Snapshot struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Files Files  `json:"files"`
}

type Snapshots []Snapshot

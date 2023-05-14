package models

import (
	"os"
	"path/filepath"
)

type Snapshot struct {
	// Name  string `json:"name"`
	Path  string `json:"path"`
	Files Files  `json:"files"`
}

type Snapshots []Snapshot

func (s *Snapshot) Load() error {
	err := filepath.Walk(s.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// skip patern
			{
				// skip snapshots dir (./)
				if path == s.Path {
					return nil
				}
				// var reg = regexp.MustCompile(`@Recently-Snapshot|@Recycle`)
				// if reg.MatchString(info.Name()) {
				// 	if info.IsDir() {
				// 		return filepath.SkipDir
				// 	}
				// 	return nil
				// }
			}
			s.Files = append(s.Files, File{Path: path})
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func (s *Snapshot) LoadFiles() error {
	for i := range s.Files {
		err := s.Files[i].Load()
		if err != nil {
			return err
		}
	}
	return nil
}

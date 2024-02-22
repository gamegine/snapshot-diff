package models

import (
	"encoding/json"

	"os"
	"path"
	"path/filepath"
	"strings"
)

type Snapshot struct {
	// Name  string `json:"name"`
	Path  string `json:"path"`
	Files Files  `json:"files"`
}

type Snapshots map[string]Snapshot

func (s *Snapshot) LoadFiles() error {
	if len(s.Files) != 0 { // not load already loaded
		return nil
	}
	resolvePath, err := filepath.EvalSymlinks(s.Path)
	if err != nil {
		return err
	}
	err = filepath.Walk(resolvePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// skip patern
			{
				// skip snapshots dir (./)
				if path == resolvePath {
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
			f := File{Path: path}
			f.LoadFileInfo(info)
			s.Files = append(s.Files, f)
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func (s *Snapshot) LoadFilesInfo() error {
	for i := range s.Files {
		err := s.Files[i].Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Snapshot) CacheFilePath(cacheDir string) string {
	return path.Join(cacheDir, s.Path[strings.LastIndex(s.Path, "/")+1:]) + ".json"
}

func (s *Snapshot) SaveCache(path string) error {
	// jsonByte, err := json.Marshal(s)
	jsonByte, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		// Error during Marshal()
		return err
	}
	err = os.WriteFile(path, jsonByte, 0644) // 0644 Unix permission bits
	if err != nil {
		// Error during WriteFile()
		return err
	}
	return nil
}

func (s *Snapshot) LoadCache(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		// Error during ReadFile
		return err
	}
	// unmarshal
	err = json.Unmarshal(content, &s)
	if err != nil {
		// Error during Unmarshal
		return err
	}
	return nil
}

func (s *Snapshot) LoadCacheOrFiles(path string) error {
	err := s.LoadCache(path)
	if err == nil {
		return nil
	}
	return s.LoadFiles()
}

func (s *Snapshot) Name() string {
	str := strings.Split(s.Path, "/")
	return str[len(str)-1]
}

func (s *Snapshot) ToMap() map[string]File {
	files := make(map[string]File)
	for _, f := range s.Files {
		files[f.Path] = f
	}
	return files
}

func (s *Snapshot) IsHash() bool {
	for _, f := range s.Files {
		if f.Sha256 == "" {
			if !IsSpecialFile(f) {
				return false
			}
		}
	}
	return true
}

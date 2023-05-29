package models

import (
	"encoding/json"
	"io/ioutil"
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

type Snapshots []Snapshot

func (s *Snapshot) Load() error {
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
			s.Files = append(s.Files, File{Path: path})
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

var SnapshotsCachePath = "./cache"

func (s *Snapshot) CachePath(vol string) string {
	return path.Join(SnapshotsCachePath, vol, s.Path[strings.LastIndex(s.Path, "/")+1:]) + ".json"
}

func (s *Snapshot) SaveCache(path string) error {
	// jsonByte, err := json.Marshal(s)
	jsonByte, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		// Error during Marshal()
		return err
	}
	err = ioutil.WriteFile(path, jsonByte, 0644) // 0644 Unix permission bits
	if err != nil {
		// Error during WriteFile()
		return err
	}
	return nil
}

func (s *Snapshot) LoadCache(path string) error {
	content, err := ioutil.ReadFile(path)
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

func (s *Snapshot) Name() string {
	str := strings.Split(s.Path, "/")
	return str[len(str)-1]
}

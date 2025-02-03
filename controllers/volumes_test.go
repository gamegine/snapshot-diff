package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"snapshot-diff/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func JSONString(j any) string {
	b, _ := json.Marshal(j)
	return string(b)
}

func TestControllerGetVolumes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	tests := []struct {
		name         string
		path         string
		Expected     string
		ExpectedCode int
	}{
		{
			name:         "empty volumes",
			path:         "./undef",
			Expected:     "{\"code\":\"500\",\"msg\":\"LoadVolumes error: open ./undef: no such file or directory\"}",
			ExpectedCode: 500,
		},
		{
			name:         "volumes",
			path:         "../testdata",
			Expected:     "{\"volumes\":[\"volume\"]}",
			ExpectedCode: 200,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			models.SnapshotsPath = TestCase.path
			r := gin.New()
			r.GET("/volumes", GetVolumes)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/volumes", nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}
		})
	}
}

func TestControllerGetVolume(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	tests := []struct {
		name         string
		path         string
		get          string
		Expected     string
		ExpectedCode int
	}{
		{
			name:         "get volume",
			get:          "volume",
			path:         "../testdata",
			Expected:     "{\"Snapshots\":[\"snapshot\",\"symlink\"],\"SnapshotsPath\":\"../testdata/volume\"}",
			ExpectedCode: 200,
		},
		{
			name:         "get undef volume",
			get:          "undef",
			path:         "./undef",
			Expected:     "{\"code\":\"500\",\"msg\":\"UpdateSnapshotsList error: open undef/undef: no such file or directory\"}",
			ExpectedCode: 500,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			models.SnapshotsPath = TestCase.path
			r := gin.New()
			r.GET("/volumes/:volume", GetVolume)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/volumes/"+TestCase.get, nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}
		})
	}
}

func TestControllerGetSnapshot(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	var s = models.Snapshot{Path: "../testdata/volume/snapshot"}
	s.LoadFiles()
	tests := []struct {
		name         string
		path         string
		get          string
		Expected     string
		ExpectedCode int
	}{
		{
			name:         "get snapshot",
			path:         "../testdata",
			get:          "volume/snapshot",
			Expected:     JSONString(s),
			ExpectedCode: 200,
		},
		{
			name:         "get undef volume",
			get:          "undef/snapshot",
			Expected:     "{\"code\":\"500\",\"msg\":\"LoadCacheOrFiles error: lstat undef: no such file or directory\"}",
			ExpectedCode: 500,
		},
		{
			name:         "get undef snapshot",
			get:          "volume/undef",
			Expected:     "{\"code\":\"500\",\"msg\":\"LoadCacheOrFiles error: lstat volume: no such file or directory\"}",
			ExpectedCode: 500,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			r := gin.New()
			models.SnapshotsPath = TestCase.path
			models.SnapshotsCachePath = "../cache"
			r.GET("/volumes/:volume/:snapshot", GetSnapshot)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/volumes/"+TestCase.get, nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}
		})
	}
}

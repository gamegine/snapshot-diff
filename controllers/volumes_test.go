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
		Volumes      models.Volumes
		Expected     string
		ExpectedCode int
	}{
		{
			name:         "empty volumes",
			Volumes:      models.Volumes{},
			Expected:     "{\"volumes\":[]}",
			ExpectedCode: 200,
		},
		{
			name:         "volumes[volume]",
			Volumes:      models.Volumes{"volume": models.Volume{}},
			Expected:     "{\"volumes\":[\"volume\"]}",
			ExpectedCode: 200,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			Volumes = TestCase.Volumes
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
		Volumes      models.Volumes
		get          string
		Expected     string
		ExpectedCode int
	}{
		{
			name:         "get volume",
			get:          "volume",
			Volumes:      models.Volumes{"volume": models.Volume{SnapshotsPath: "SnapshotsPath"}},
			Expected:     "{\"Snapshots\":[],\"SnapshotsPath\":\"SnapshotsPath\"}",
			ExpectedCode: 200,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			Volumes = TestCase.Volumes
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
	var s = models.Snapshot{Path: "undefined"} // path undefined for not LoadFiles
	tests := []struct {
		name         string
		Volumes      models.Volumes
		get          string
		Expected     string
		ExpectedCode int
	}{
		{
			name: "get snapshot",
			get:  "volume/snapshot",
			Volumes: models.Volumes{"volume": models.Volume{
				Snapshots: models.Snapshots{"snapshot": s},
			},
			},
			Expected:     JSONString(s),
			ExpectedCode: 200,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			Volumes = TestCase.Volumes
			r := gin.New()
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

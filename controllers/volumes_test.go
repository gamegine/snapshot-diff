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

func TestInitVolumes(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name     string
		path     string
		Expected models.Volumes
		error    bool
	}{
		{
			name: "volumes",
			path: "../testdata",
			Expected: models.Volumes{"volume": models.Volume{
				SnapshotsPath: "../testdata/volume",
				Snapshots: models.Snapshots{"snapshot": {
					Path:        "testdata/volume/snapshot",
					ResolvePath: "testdata/volume/snapshot",
				}},
			}},
			error: false,
		},
		{
			name:     "error snapshots path",
			path:     "/error",
			Expected: models.Volumes{},
			error:    true,
		},
	}
	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			models.SnapshotsCachePath = "../cache"
			models.SnapshotsPath = TestCase.path
			err := InitVolumes()
			if TestCase.error {
				if err == nil {
					t.Errorf("error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("error %v", err)
				}
			}
		})
	}
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
		{
			name:         "get undef volume",
			get:          "undef",
			Volumes:      models.Volumes{},
			Expected:     "{\"code\":\"404\",\"msg\":\"volume not found\"}",
			ExpectedCode: 404,
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
				Snapshots: models.Snapshots{"snapshot": s}},
			},
			Expected:     JSONString(s),
			ExpectedCode: 200,
		},
		{
			name:         "get undef volume",
			get:          "undef/snapshot",
			Volumes:      models.Volumes{},
			Expected:     "{\"code\":\"404\",\"msg\":\"volume not found\"}",
			ExpectedCode: 404,
		},
		{
			name: "get undef snapshot",
			get:  "volume/undef",
			Volumes: models.Volumes{"volume": models.Volume{
				Snapshots: models.Snapshots{"snapshot": s}},
			},
			Expected:     "{\"code\":\"404\",\"msg\":\"snapshot not found\"}",
			ExpectedCode: 404,
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

func TestControllerUpdateVolumes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	tests := []struct {
		name            string
		Volumes         models.Volumes
		path            string
		Expected        string
		ExpectedCode    int
		ExpectedVolumes int
	}{
		{
			name:            "update",
			path:            "../testdata",
			Volumes:         models.Volumes{},
			Expected:        "{\"code\":\"200\",\"msg\":\"ok\"}",
			ExpectedCode:    200,
			ExpectedVolumes: 1,
		},
		{
			name:            "error",
			path:            "/error",
			Volumes:         models.Volumes{},
			Expected:        "{\"code\":\"500\",\"msg\":{\"Op\":\"open\",\"Path\":\"/error\",\"Err\":2}}",
			ExpectedCode:    500,
			ExpectedVolumes: 0,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {

			models.SnapshotsCachePath = "../cache"
			models.SnapshotsPath = TestCase.path

			Volumes = TestCase.Volumes
			r := gin.New()
			r.POST("/volumes", UpdateVolumes)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/volumes", nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}
			if TestCase.ExpectedVolumes != len(Volumes) {
				t.Errorf("got %v, wanted %v", TestCase.ExpectedVolumes, len(Volumes))
			}
		})
	}
}

func TestControllerUpdateVolume(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	tests := []struct {
		name              string
		Volumes           models.Volumes
		get               string
		Expected          string
		ExpectedCode      int
		ExpectedSnapshots []string
	}{
		{
			name:              "update",
			get:               "volume",
			Volumes:           models.Volumes{"volume": models.Volume{SnapshotsPath: "../testdata/volume"}},
			Expected:          "{\"code\":\"200\",\"msg\":\"ok\"}",
			ExpectedCode:      200,
			ExpectedSnapshots: []string{"snapshot", "symlink"},
		},
		{
			name:              "undef volume",
			get:               "undef",
			Volumes:           models.Volumes{"volume": models.Volume{SnapshotsPath: "../testdata/volume"}},
			Expected:          "{\"code\":\"404\",\"msg\":\"volume not found\"}",
			ExpectedCode:      404,
			ExpectedSnapshots: []string{},
		},
		{
			name:              "error",
			get:               "volume",
			Volumes:           models.Volumes{"volume": models.Volume{SnapshotsPath: "./undef"}},
			Expected:          "{\"code\":\"500\",\"msg\":{\"Op\":\"open\",\"Path\":\"./undef\",\"Err\":2}}",
			ExpectedCode:      500,
			ExpectedSnapshots: []string{},
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			Volumes = TestCase.Volumes
			r := gin.New()
			r.POST("/volume/:volume", UpdateVolume)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/volume/"+TestCase.get, nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}

			v := Volumes["volume"]
			if len(TestCase.ExpectedSnapshots) != len(v.Snapshots) {
				t.Errorf("got %v, wanted %v", len(TestCase.ExpectedSnapshots), len(v.Snapshots))
			}
		})
	}
}

func TestControllerUpdateSnapshot(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Anonymous struct of test cases
	tests := []struct {
		name          string
		Volumes       models.Volumes
		get           string
		Expected      string
		ExpectedCode  int
		ExpectedFiles int
	}{
		{
			name: "update",
			get:  "volume/snapshot",
			Volumes: models.Volumes{"volume": models.Volume{
				Snapshots: models.Snapshots{"snapshot": models.Snapshot{Path: "../testdata/volume"}}},
			},
			Expected:      "{\"code\":\"200\",\"msg\":\"ok\"}",
			ExpectedCode:  200,
			ExpectedFiles: 4,
		},
		{
			name:          "undef volume",
			get:           "undef/snapshot",
			Volumes:       models.Volumes{},
			Expected:      "{\"code\":\"404\",\"msg\":\"volume not found\"}",
			ExpectedCode:  404,
			ExpectedFiles: 0,
		},
		{
			name: "undef snapshot",
			get:  "volume/undef",
			Volumes: models.Volumes{"volume": models.Volume{
				Snapshots: models.Snapshots{"snapshot": models.Snapshot{Path: "../testdata/volume"}}},
			},
			Expected:      "{\"code\":\"404\",\"msg\":\"snapshot not found\"}",
			ExpectedCode:  404,
			ExpectedFiles: 0,
		},
		{
			name: "error",
			get:  "volume/snapshot",
			Volumes: models.Volumes{"volume": models.Volume{
				Snapshots: models.Snapshots{"snapshot": models.Snapshot{Path: "undefined"}}},
			},
			Expected:      "{\"code\":\"500\",\"msg\":{\"Op\":\"lstat\",\"Path\":\"undefined\",\"Err\":2}}",
			ExpectedCode:  500,
			ExpectedFiles: 0,
		},
	}

	for _, TestCase := range tests {
		// each test case from  table above run as a subtest
		t.Run(TestCase.name, func(t *testing.T) {
			Volumes = TestCase.Volumes
			r := gin.New()
			r.POST("/volumes/:volume/:snapshot", UpdateSnapshot)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/volumes/"+TestCase.get, nil)
			r.ServeHTTP(w, req)

			if w.Code != TestCase.ExpectedCode {
				t.Errorf("got %v, wanted %v", w.Code, TestCase.ExpectedCode)
			}
			if w.Body.String() != TestCase.Expected {
				t.Errorf("got %v, wanted %v", w.Body.String(), TestCase.Expected)
			}
			if len(Volumes["volume"].Snapshots["snapshot"].Files) != TestCase.ExpectedFiles {
				t.Errorf("got %v, wanted %v", len(Volumes["volume"].Snapshots["snapshot"].Files), TestCase.ExpectedFiles)
			}
		})
	}
}

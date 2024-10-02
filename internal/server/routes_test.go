package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mishakrpv/musiclib/internal/endpoint/command"
	"github.com/mishakrpv/musiclib/internal/mock"

	"github.com/gin-gonic/gin"
)

func TestCreateSongHandler(t *testing.T) {
	ConfigureLogging()
	testServer := &Server{
		songRepo: mock.NewSongRepository(),

		musicInfoClient: mock.NewGoodMusicInfoClient(),
	}
	r := gin.New()
	r.POST("/songs", testServer.CreateSongHandler)

	var tests = []struct {
		name           string
		expectedStatus int
		group          string
		song           string
	}{
		{"Correct inputs should be 200", http.StatusOK, "Muse", "Supermassive Black Hole"},
		{"No song should be 400", http.StatusBadRequest, "Muse", ""},
	}

	for _, test := range tests {
		t.Logf("Start running test: %s", test.name)
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			request := &command.CreateRequest{
				Group: test.group,
				Song:  test.song,
			}

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(request)
			if err != nil {
				t.Fatal(err)
			}

			// Act
			req, err := http.NewRequest("POST", "/songs", &buf)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			r.ServeHTTP(recorder, req)

			// Assert
			if status := recorder.Code; status != test.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, test.expectedStatus)
			}
		})
	}
}

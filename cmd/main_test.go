package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	tracker "github.com/reactivejson/flaypath/internal"
)

func TestGetFlightPath(t *testing.T) {
	router := gin.Default()
	router.POST("/calculate", GetFlightPath)

	testCases := []struct {
		name     string
		input    []tracker.Flight
		expected []string
	}{
		{
			name:     "single-flight",
			input:    []tracker.Flight{{Source: "SFO", Destination: "EWR"}},
			expected: []string{"SFO", "EWR"},
		},
		{
			name:     "multiple-flights",
			input:    []tracker.Flight{{Source: "ATL", Destination: "EWR"}, {Source: "SFO", Destination: "ATL"}},
			expected: []string{"SFO", "EWR"},
		},
		{
			name: "multiple-flights-with-layover",
			input: []tracker.Flight{
				{Source: "IND", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
			},
			expected: []string{"SFO", "EWR"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest("POST", "/calculate", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			var result map[string][]string
			json.Unmarshal(resp.Body.Bytes(), &result)
			//TODO assert.NoError(t, err)

			//TODO assert.Equal(t, tc.expected, result["calculate"])
		})
	}

}

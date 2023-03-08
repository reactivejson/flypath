package api

import (
	"encoding/json"
	"github.com/reactivejson/flaypath/.idea/internal/flights"
	"net/http"

	"github.com/gorilla/mux"
)

type request struct {
	Flights [][]string `json:"flights"`
}

type response struct {
	Path []string `json:"path"`
}

// FindPathHandler handles the /find-path endpoint
func FindPathHandler(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path := flights.FindPath(req.Flights)

	resp := response{Path: path}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterHandlers registers the API handlers with the given router
func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/find-path", FindPathHandler).Methods(http.MethodPost)
}

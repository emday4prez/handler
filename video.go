package main

import (
	"encoding/json"
	"net/http"
)

type Video struct {
	ID    string
	Title string
}

func (s *Server) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	var video Video

	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {

		http.Error(w, "bad request: invalid JSON", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"status":      "success",
		"id_received": video.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

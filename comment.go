package main

import (
	"encoding/json"
	"net/http"
)

type Comment struct {
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "error decoding json", http.StatusBadRequest)
		return
	}

	if comment.Text == "" {
		http.Error(w, "error - there is no comment text", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(comment)

}

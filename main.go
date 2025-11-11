package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	videoId := r.URL.Query().Get("video_id")
	if videoId == "" {
		http.Error(w, "no video id param", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Fetching comments for video " + videoId))

}

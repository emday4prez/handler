package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetCommentHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/comment?video_id=3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	GetCommentHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status: %v (not 200)", status)
	}

	if body := rr.Body.String(); !strings.Contains(body, "v123") {
		t.Errorf("body doesnt contain v123")
	}
}

func TestPostCommentHandler(t *testing.T) {

	mux := chi.NewRouter()
	mux.Post("/video/{videoID}/comment", PostCommentHandler)
	payload := `{"user_id": "u352", "text": "this is such a comment"}`

	req := httptest.NewRequest("POST", "/video/v789/comment", strings.NewReader(payload))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}
	if !strings.Contains(rr.Body.String(), "v789") {
		t.Errorf("handler body did not include videoID: got %v", rr.Body.String())
	}
}

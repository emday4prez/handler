package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
	payload := `{"user_id": "u352", "text": "this is such a comment"}`

	req := httptest.NewRequest("POST", "/comments", strings.NewReader(payload))

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	PostCommentHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code")
	}
}

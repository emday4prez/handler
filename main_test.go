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
	handler := http.HandlerFunc(GetCommentHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status: %v (not 200)", status)
	}

	if body := rr.Body.String(); !strings.Contains(body, "v123") {
		t.Errorf("body doesnt contain v123")
	}
}

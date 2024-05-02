package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeStaticFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/testfile.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	serveStaticFile(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Error("handler returned valid status code ",
			status, http.StatusOK)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}
}

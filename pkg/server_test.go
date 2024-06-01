package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeStaticFile(t *testing.T) {
	config = MockConfig()
	config.Routes = []string{"/index"}
	config.FallbackDocument = "index.html"
	config.RootDir = "index.html"
	config.Port = "3000"

	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveStaticFile)

	testDataDir := filepath.Join("..", "public")
	err = os.MkdirAll(testDataDir, 0755)

	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testDataDir)

	indexFile := filepath.Join(testDataDir, "index.html")

	err = os.WriteFile(indexFile, []byte("Hello, world!"), 0644)

	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// if rr.Body.String() != "Hello, world!" {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), "Hello, world!")
	// }

}

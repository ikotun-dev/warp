package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	config := InitConfig()

	filePath := "./examples/static/" + r.URL.Path
	_, err := os.Stat(filePath)

	if err != nil {
		fallbackDocument := config.FallbackDocument
		filePath = filepath.Join("./examples/static/" + fallbackDocument)
	}

}

func server() {
	fs := http.FileServer(http.Dir("./examples/static"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

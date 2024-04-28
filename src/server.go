package main

import (
	"net/http"
	"os"
)

func serveStaticFile(w http.ResponseWriter, r *http.Request) {

	filePath := "./examples/static/" + r.URL.Path
	_, err := os.Stat(filePath)
	if err != nil {

	}

}

func server() {
	fs := http.FileServer(http.Dir("./examples/static"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

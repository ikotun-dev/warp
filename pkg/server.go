package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var config *Config
var mimeType string
var filePath string

func serveStaticFile(w http.ResponseWriter, r *http.Request) {

	if config == nil {
		http.Error(w, "Configuration not initialized", http.StatusInternalServerError)
		return
	}

	filePath = "../public" + r.URL.Path

	_, err := os.Stat(filePath)
	defaultPath := filepath.Join("../public", config.RootDir)

	if err != nil {
		filePath = config.RootDir
		log.Printf("Rewriting to root file :  %s", r.URL.Path)
		w.WriteHeader(http.StatusTemporaryRedirect)
		http.ServeFile(w, r, defaultPath)
	}

	ext := filepath.Ext(filePath)

	if ext != "" {
		mimeType = getMimeType(ext)
	} else {
		mimeType = "text/html; charset=utf-8"
	}

	w.Header().Set("Content-Type", mimeType)

	http.ServeFile(w, r, filePath)
}

// Entry
// Initiates the config.yaml file
func main() {
	config = InitConfig()

	http.HandleFunc("/", serveStaticFile)

	fmt.Println("Warp server running on : ", config.Port)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

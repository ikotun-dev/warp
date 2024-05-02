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
	//TODO: handle rewrites, 404

	if config == nil {
		http.Error(w, "Configuration not initialized", http.StatusInternalServerError)
		return
	}

	//TODO: Loggin option in the config yaml

	filePath = "../public" + r.URL.Path

	_, err := os.Stat(filePath)
	defaultPath := filepath.Join("../public", config.RootDir)

	if err != nil {
		//TRYING TO SEE WHAT I CAN DO FOR THE 404 STUFF
		//
		filePath = config.RootDir
		log.Printf("LOG: I no see the file : %s", r.URL.Path)
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

	configPort := ":" + config.Port

	fmt.Println("PORT : ", configPort)

	log.Fatal(http.ListenAndServe(configPort, nil))
}

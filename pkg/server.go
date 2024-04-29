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
	//Make heavy loggin a config option

	filePath = "./examples/static" + r.URL.Path
	_, err := os.Stat(filePath)

	if err != nil {

		fallbackDocument := config.FallbackDocument
		filePath = filepath.Join("./examples/static", fallbackDocument)
	}

	ext := filepath.Ext(filePath)

	if ext != "" {
		mimeType = getMimeType(ext)
	} else {
		mimeType = "text/html; charset=utf-8"
	}

	log.Printf("Extension: %s", ext)

	log.Printf("MimeType :  %s", mimeType)

	w.Header().Set("Content-Type", mimeType)

	fmt.Println("Server running on  : ", config.Port)
	http.ServeFile(w, r, filePath)
}

func main() {
	config = InitConfig()
	http.HandleFunc("/", serveStaticFile)

	configPort := ":" + config.Port

	fmt.Println("PORT : ", configPort)

	err := http.ListenAndServe(configPort, nil)

	if err != nil {
		log.Fatal("Error : ", err)
	}
}
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var config *Config

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	if config == nil {
		http.Error(w, "Configuration not initialized", http.StatusInternalServerError)
		return
	}

	filePath := "./examples/static" + r.URL.Path
	_, err := os.Stat(filePath)

	if err != nil {
		fallbackDocument := config.FallbackDocument
		filePath = filepath.Join("./examples/static", fallbackDocument)
	}

	ext := filepath.Ext(filePath)

	mimeType := getMimeType(ext[1:])

	fmt.Println("MimeType : ", mimeType)

	w.Header().Set("Content-Type", mimeType)

	fmt.Println("Server running on  : %s", config.Port)
	http.ServeFile(w, r, filePath)
}

func main() {
	config = InitConfig()
	http.HandleFunc("/", serveStaticFile)

	configPort := ":" + config.Port

	fmt.Println("PORT : ", configPort)

	http.ListenAndServe(configPort, nil)
}

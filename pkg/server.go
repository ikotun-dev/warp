//TODO: follow conventions [www/var/html]

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

	//TODO: check if the path isnt in the rewrite
	//
	filePath = "../public" + r.URL.Path

	_, err := os.Stat(filePath)
	defaultPath := filepath.Join("../public", config.RootDir)

	if err != nil {

		// check if the route is in the list of provided routes.

		eixstingRoute := false
		for _, route := range config.Routes {
			log.Printf("Route : %s", route)
			if r.URL.Path == route {
				eixstingRoute = true
			}
		}

		// filePath = config.RootDir
		// log.Printf("Rewriting to root file :  %s", r.URL.Path)
		// w.WriteHeader(http.StatusTemporaryRedirect)

		if eixstingRoute {
			http.ServeFile(w, r, defaultPath)
		} else {
			if config.FallbackDocument != "" {
				fallBackDocumentPath := filepath.Join("../public", config.FallbackDocument)
				http.ServeFile(w, r, fallBackDocumentPath)

			}

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 Not Found, served by Warp Server")

		}
	}

	ext := filepath.Ext(filePath)

	if ext != "" {
		mimeType = getMimeType(ext)
	} else {
		mimeType = "text/html; charset=utf-8"
	}

	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
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

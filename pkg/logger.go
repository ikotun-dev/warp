package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

		// Call the next handler
		handler.ServeHTTP(w, r)
	})
}

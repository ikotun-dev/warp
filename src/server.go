package main

import (
	"net/http"
)

func server() {
	fs := http.FileServer(http.Dir("./examples/static"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

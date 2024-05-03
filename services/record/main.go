package main

import (
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

// Package httpserver ...
package httpserver

import (
	"io"
	"net/http"
)

// StringHandler ...
type StringHandler struct {
	message string
}

// ServeHTTP method realized
func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - returning 404")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", r.URL.Path)
	io.WriteString(w, sh.message)
}

// Start ...
func Start() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World!"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

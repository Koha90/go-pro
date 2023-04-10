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
	Printfln("Request for %v", r.URL.Path)
	switch r.URL.Path {
	case "/favicon.ico":
		http.NotFound(w, r)
	case "/message":
		io.WriteString(w, sh.message)
	default:
		http.Redirect(w, r, "/message", http.StatusTemporaryRedirect)
	}
}

// Start ...
func Start() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World!"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

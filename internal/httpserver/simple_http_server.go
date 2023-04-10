// Package httpserver ...
package httpserver

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, sh.message)
}

func Start() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World!"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

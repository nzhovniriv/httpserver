package web

import (
	"fmt"
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got %s request\n", r.URL.Path)
	io.WriteString(w, "Hello, HTTP!\n")
}

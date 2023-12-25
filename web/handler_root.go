package web

import (
	"fmt"
	"io"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got %s request\n", r.URL.Path)
	io.WriteString(w, "This is my website!\n")
}

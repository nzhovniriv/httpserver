package main

import (
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"httpserver/web"
)

func main() {
	http.HandleFunc("/hello", web.Hello)
	http.HandleFunc("/", web.Root)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		}
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// TCP链接，handler参数是nil表示使用DefaultServeMux
	http.ListenAndServe(":80", nil)
}

package main

import (
	"net/http"
)

func main() {
	handler, _ := QueryBlogsHandler()

	http.Handle("/blogs", handler)

	http.ListenAndServe(":8000", nil)
}

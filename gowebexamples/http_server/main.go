package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// the http.stripprefix removes the /static/ prefix from the request URL,
	// so that the fs gets only file name not the whole path
	// as fs is going to look for files in static folder so we need only filename not the whole path.
	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080")
}

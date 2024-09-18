package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", loggerDup)
	http.HandleFunc("/bar", logger(bar))

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}

func loggerDup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log")
	foo(w, r)
}

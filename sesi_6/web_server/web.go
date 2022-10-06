package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hi there !"
	fmt.Fprint(w, msg)
}

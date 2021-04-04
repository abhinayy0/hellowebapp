package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World! from Golang.")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}
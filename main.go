package main

import (
	"fmt"
	"net/http"
)

func helloWordHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func goodByeWorldHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GoodBye World")
}

func main() {

	http.HandleFunc("/", helloWordHandler)
	http.HandleFunc("/goodbye", goodByeWorldHandler)
	http.ListenAndServe(":8080", nil)
}

// localhost:8080
//fetch(localhost:8080)
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Unsupported method call", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func formHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}

	name := r.PostFormValue("name")
	address := r.PostFormValue("address")

	fmt.Fprintf(w, "Name: %s and Address: %s", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandlerFunc)
	http.HandleFunc("/form", formHandlerFunc)

	fmt.Printf("Starting server at port http://localhost:8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

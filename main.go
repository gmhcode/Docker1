package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world f")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8081", nil)

	fmt.Println("hello")
}

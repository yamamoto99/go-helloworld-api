package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Hello World"}`)
}

func main() {
	http.HandleFunc("/hello_world", helloWorld)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

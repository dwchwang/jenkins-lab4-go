package main

import (
	"fmt"
	"net/http"
)

func add(a, b int) int {
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go! 2+3=%d", add(2, 3))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}

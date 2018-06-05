package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Hello, WWW!</H1>")
}

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println(http.ListenAndServe(":8080", nil)) // HL
}

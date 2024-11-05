package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	fmt.Printf("Server running on %s", port)
	http.ListenAndServe(port, http.FileServer(http.Dir(".")))
}

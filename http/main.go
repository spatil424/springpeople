package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Welcome to go http")
	})
	fmt.Println("server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}

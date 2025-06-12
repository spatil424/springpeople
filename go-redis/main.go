package main

import (
	"fmt"
	"goredis/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi")
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

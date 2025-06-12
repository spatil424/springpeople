package main

import (
	"fmt"
	"goRedis-advanced/handlers"
	redisclient "goRedis-advanced/redis"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server running on http://localhost:4000")
	redisclient.Init()
	log.Fatal(http.ListenAndServe(":4000", nil))
}

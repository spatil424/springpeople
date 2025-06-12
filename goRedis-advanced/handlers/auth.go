package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	rdb "goRedis-advanced/redis"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	creds.Username = strings.ToLower(creds.Username)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	key := "user:" + creds.Username
	rdb.Rdb.HSet(rdb.Ctx, key, "password", hashedPassword)
	rdb.Rdb.SAdd(rdb.Ctx, "users", creds.Username)
	w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	creds.Username = strings.ToLower(creds.Username)

	key := "user:" + creds.Username
	storedHash, err := rdb.Rdb.HGet(rdb.Ctx, key, "password").Result()
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Login successful"))
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := rdb.Rdb.SMembers(rdb.Ctx, "users").Result()
	json.NewEncoder(w).Encode(users)
}

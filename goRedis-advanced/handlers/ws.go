package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	rdb "goRedis-advanced/redis"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[string]*websocket.Conn)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	clients[username] = conn

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, username)
			return
		}

		chatKey := chatKey(msg.From, msg.To)
		jsonMsg, _ := json.Marshal(msg)
		rdb.Rdb.RPush(rdb.Ctx, chatKey, jsonMsg)

		if receiver, ok := clients[msg.To]; ok {
			receiver.WriteJSON(msg)
		}
	}
}

func chatKey(a, b string) string {
	users := []string{a, b}
	sort.Strings(users)
	return "chat:" + users[0] + ":" + users[1]
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", handleWebSocket)
	http.Handle("/", r)
	log.Println("Server is running: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Пример отправки SDUI-контента по WebSocket
	response := []byte(`{
		"components": [
			{
				"type": "text",
				"content": "Welcome to our app!"
			},
			{
				"type": "image",
				"url": "https://aif-s3.aif.ru/images/018/114/20868fd6f3f1725ed7ff96806255e610.jpg"
			},
			{
				"type": "button",
				"label": "Click me",
				"action": "open_link",
				"data": "https://minecraft.net"
			}
		]
	}`)

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}

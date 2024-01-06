package ws

import "github.com/gorilla/websocket"

type Client struct {
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
	Conn     *websocket.Conn
	Message  chan *Message
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

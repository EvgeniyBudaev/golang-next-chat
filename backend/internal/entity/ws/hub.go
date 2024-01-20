package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/searching"
)

type Room struct {
	ID       int64  `json:"id"`
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}

type RoomProfile struct {
	ID        int64 `json:"id"`
	RoomID    int64 `json:"roomId"`
	ProfileID int64 `json:"profileId"`
}

type RoomWithProfileResponse struct {
	ID       int64  `json:"id"`
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}

type QueryParamsRoomList struct {
	searching.Searching
}

type QueryParamsJoinRoom struct {
	ClientID string `json:"clientId"`
	Username string `json:"username"`
}

type Hub struct {
	Clients    map[int64][]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Content
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64][]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Content, 5),
	}
}

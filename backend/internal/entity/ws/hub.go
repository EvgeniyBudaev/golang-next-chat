package ws

import (
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/searching"
	"github.com/google/uuid"
)

type Room struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	RoomName string    `json:"roomName"`
	Title    string    `json:"title"`
}

type RoomProfile struct {
	ID        int64 `json:"id"`
	RoomID    int64 `json:"roomId"`
	ProfileID int64 `json:"profileId"`
}

type RoomWithProfileResponse struct {
	ID       int64                                 `json:"id"`
	UUID     uuid.UUID                             `json:"room"`
	RoomName string                                `json:"roomName"`
	Title    string                                `json:"title"`
	Profile  *profileEntity.ResponseProfileForRoom `json:"profile"`
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
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64][]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

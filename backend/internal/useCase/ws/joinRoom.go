package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/gofiber/contrib/websocket"
)

type JoinRoomUseCase struct {
	hub *ws.Hub
}

func NewJoinRoomUseCase(h *ws.Hub) *JoinRoomUseCase {
	return &JoinRoomUseCase{
		hub: h,
	}
}

func (uc *JoinRoomUseCase) JoinRoom(conn *websocket.Conn) string {
	roomId := conn.Params("roomId")
	clientId := conn.Query("clientId")
	username := conn.Query("username")
	cl := &ws.Client{
		ID:       clientId,
		RoomID:   roomId,
		Username: username,
		Conn:     conn,
		Message:  make(chan *ws.Message),
	}
	m := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomId,
		Username: username,
	}
	// Register a new client through the register channel
	uc.hub.Register <- cl
	// Broadcast that message
	uc.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return "is joined"
}

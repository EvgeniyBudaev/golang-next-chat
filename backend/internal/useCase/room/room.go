package room

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type UseCaseRoom struct {
	hub *ws.Hub
}

func NewUseCaseRoom(h *ws.Hub) *UseCaseRoom {
	return &UseCaseRoom{
		hub: h,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (uc *UseCaseRoom) CreateRoom(ctx *fiber.Ctx, r CreateRoomRequest) (*ws.RoomResponse, error) {
	// Хранение информации по комнате в памяти, а не в БД
	uc.hub.Rooms[r.ID] = &ws.Room{
		ID:      r.ID,
		Name:    r.Name,
		Clients: make(map[string]*ws.Client),
	}
	response := ws.RoomResponse{
		ID:   r.ID,
		Name: r.Name,
	}
	//ctx.Status(fiber.StatusCreated).JSON(r)
	return &response, nil
}

func (uc *UseCaseRoom) GetClientList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error) {
	clients := make([]*ws.ClientResponse, 0)
	roomId := ctx.Params("roomId")
	if _, ok := uc.hub.Rooms[roomId]; !ok {
		return clients, nil
	}
	for _, c := range uc.hub.Rooms[roomId].Clients {
		clients = append(clients, &ws.ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}
	return clients, nil
}

func (uc *UseCaseRoom) GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomResponse, error) {
	rooms := make([]*ws.RoomResponse, 0)
	for _, r := range uc.hub.Rooms {
		rooms = append(rooms, &ws.RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	return rooms, nil
}

func (uc *UseCaseRoom) JoinRoom(conn *websocket.Conn) string {
	roomId := conn.Params("roomId")
	clientId := conn.Query("userId")
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
	// Register a new client through the user channel
	uc.hub.Register <- cl
	// Broadcast that message
	uc.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return "is joined"
}

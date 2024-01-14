package room

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/room"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type UseCaseRoom struct {
	hub *ws.Hub
	db  *room.PGRoomDB
}

func NewUseCaseRoom(h *ws.Hub, db *room.PGRoomDB) *UseCaseRoom {
	return &UseCaseRoom{
		hub: h,
		db:  db,
	}
}

type CreateRoomRequest struct {
	Name string `json:"name"`
}

func (uc *UseCaseRoom) Run(ctx *fiber.Ctx) {
	for {
		select {
		case cl := <-uc.hub.Register:
			fmt.Println("hub Register: ", cl)
			//if _, ok := h.Rooms[cl.RoomID]; ok {
			//	r := h.Rooms[cl.RoomID]
			//
			//	if _, ok := r.Clients[cl.ID]; !ok {
			//		r.Clients[cl.ID] = cl
			//	}
			//}
		case cl := <-uc.hub.Unregister:
			fmt.Println("hub Unregister: ", cl)
			//if _, ok := h.Rooms[cl.RoomID]; ok {
			//	if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
			//		if len(h.Rooms[cl.RoomID].Clients) != 0 {
			//			h.Broadcast <- &Message{
			//				RoomID:   cl.RoomID,
			//				ClientID: cl.ID,
			//				Content:  "user left the chat",
			//			}
			//		}
			//		delete(h.Rooms[cl.RoomID].Clients, cl.ID)
			//		close(cl.Message)
			//	}
			//}
		case m := <-uc.hub.Broadcast:
			fmt.Println("hub Broadcast: ", m)
			clientList, err := uc.db.SelectClientList()
			if err != nil {
				logger.Log.Debug("error func Run, method SelectClientList by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			for _, item := range clientList {
				_, err := uc.db.AddMessage(m)
				if err != nil {
					logger.Log.Debug("error func Run, method SelectClientList by path internal/useCase/room/room.go",
						zap.Error(err))
				}
				fmt.Println("item: ", item)
			}
			//if _, ok := h.Rooms[m.RoomID]; ok {
			//
			//	for _, cl := range h.Rooms[m.RoomID].Clients {
			//		cl.Message <- m
			//	}
			//}
		}
	}
}

func (uc *UseCaseRoom) CreateRoom(ctx *fiber.Ctx, r CreateRoomRequest) (*ws.RoomResponse, error) {
	roomRequest := &ws.Room{
		Name: r.Name,
	}
	newRoom, err := uc.db.Create(ctx, roomRequest)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method Create by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	response := ws.RoomResponse{
		ID:   newRoom.ID,
		Name: newRoom.Name,
	}
	return &response, nil
}

func (uc *UseCaseRoom) GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomResponse, error) {
	roomList, err := uc.db.SelectRoomList(ctx)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	response := make([]*ws.RoomResponse, 0)
	for _, item := range roomList {
		roomResponse := &ws.RoomResponse{
			ID:   item.ID,
			Name: item.Name,
		}
		response = append(response, roomResponse)
	}
	return response, nil
}

func (uc *UseCaseRoom) GetClientList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error) {
	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method ParseInt by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	fmt.Println("roomId: ", roomId)
	clientList, err := uc.db.SelectClientList()
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	response := make([]*ws.ClientResponse, 0)
	for _, item := range clientList {
		roomResponse := &ws.ClientResponse{
			ID:       item.ID,
			Username: item.Username,
		}
		response = append(response, roomResponse)
	}
	return response, nil
}

func (uc *UseCaseRoom) JoinRoom(conn *websocket.Conn) string {
	roomIdStr := conn.Params("roomId")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return ""
	}
	userId := conn.Query("userId")
	username := conn.Query("username")

	cl := &ws.Client{
		RoomID:   roomId,
		UserID:   userId,
		Username: username,
		Conn:     conn,
		Message:  make(chan *ws.Message),
	}
	// Register a new client through the user channel
	newClient, err := uc.db.AddClient(cl)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method AddClient by path internal/useCase/room/room.go",
			zap.Error(err))
	}
	m := &ws.Message{
		RoomID:   roomId,
		ClientID: newClient.ID,
		Content:  "A new user has joined the room",
	}
	//uc.hub.Register <- cl
	// Broadcast that message
	uc.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return newClient.Username
}

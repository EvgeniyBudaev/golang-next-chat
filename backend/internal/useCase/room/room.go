package room

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"

	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/room"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	UserID string `json:"userId"`
}

type GetRoomMessagesRequest struct {
	RoomID string `json:"roomId"`
}

func (uc *UseCaseRoom) Run(ctx *fiber.Ctx) {
	for {
		select {
		case cl := <-uc.hub.Register:
			fmt.Println("hub Register: ", cl)
			uc.hub.Clients[cl.RoomID] = append(uc.hub.Clients[cl.RoomID], cl)
			uc.hub.Broadcast <- &ws.Message{
				UUID:      uuid.New(),
				RoomID:    cl.RoomID,
				UserID:    cl.UserID,
				Type:      ws.SystemMessage,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
				IsEdited:  false,
				Content:   cl.Username + " has joined the channel",
			}

		case cl := <-uc.hub.Unregister:
			fmt.Println("hub Unregister: ", cl)
			if _, ok := uc.hub.Clients[cl.RoomID]; ok {
				for i, c := range uc.hub.Clients[cl.RoomID] {
					if c == cl {
						uc.hub.Clients[cl.RoomID] = append(uc.hub.Clients[cl.RoomID][:i], uc.hub.Clients[cl.RoomID][i+1:]...)
						break
					}
				}
			}
			uc.hub.Broadcast <- &ws.Message{
				UUID:      uuid.New(),
				RoomID:    cl.RoomID,
				UserID:    cl.UserID,
				Type:      ws.SystemMessage,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
				IsEdited:  false,
				Content:   cl.Username + " left the channel",
			}

		case m := <-uc.hub.Broadcast:
			fmt.Println("hub Broadcast: ", m)
			_, err := uc.db.AddMessage(m)
			if err != nil {
				logger.Log.Debug("error func Run, method AddMessage by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			if _, ok := uc.hub.Clients[m.RoomID]; ok {
				for _, cl := range uc.hub.Clients[m.RoomID] {
					cl.Message <- m
				}
			}
		}
	}
}

func (uc *UseCaseRoom) CreateRoom(ctx *fiber.Ctx, r CreateRoomRequest) (*ws.Room, error) {
	roomRequest := &ws.Room{
		UUID: uuid.New(),
	}
	newRoom, err := uc.db.CreateRoom(ctx, roomRequest)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method Create by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	profile, err := uc.db.FindProfile(r.UserID)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method FindProfile by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	_, err = uc.db.AddRoomProfile(newRoom.ID, profile.ID)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method AddRoomProfile by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return newRoom, nil
}

func (uc *UseCaseRoom) GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error) {
	response, err := uc.db.SelectRoomList(ctx)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return response, nil
}

func (uc *UseCaseRoom) GetUserList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error) {
	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method ParseInt by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	fmt.Println("roomId: ", roomId)
	clientList, err := uc.db.SelectUserList()
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

func (uc *UseCaseRoom) GetMessageList(ctx *fiber.Ctx, r GetRoomMessagesRequest) ([]*ws.ResponseMessage, error) {
	roomId, err := strconv.ParseInt(r.RoomID, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	messageList, err := uc.db.SelectMessageList(ctx, roomId)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return messageList, nil
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
	profile, err := uc.db.FindProfile(userId)
	_, err = uc.db.AddRoomProfile(roomId, profile.ID)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method AddUser by path internal/useCase/room/room.go",
			zap.Error(err))
	}
	uc.hub.Register <- cl
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return ""
}

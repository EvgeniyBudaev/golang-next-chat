package room

import (
	"fmt"
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
	UserID   string `json:"userId"`
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}

type GetRoomListByProfileRequest struct {
	ProfileID string `json:"profileId"`
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
			message := &ws.Message{
				RoomID:    cl.RoomID,
				UserID:    cl.UserID,
				Type:      ws.SystemMessage,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
				IsEdited:  false,
				Content:   cl.Username + " has joined the channel",
			}
			uc.hub.Broadcast <- &ws.Content{
				Message: message,
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
			message := &ws.Message{
				RoomID:    cl.RoomID,
				UserID:    cl.UserID,
				Type:      ws.SystemMessage,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
				IsEdited:  false,
				Content:   cl.Username + " left the channel",
			}
			uc.hub.Broadcast <- &ws.Content{
				Message: message,
			}

		case c := <-uc.hub.Broadcast:
			fmt.Println("hub Broadcast: ", c)
			start := time.Now()
			_, err := uc.db.AddMessage(c.Message)
			elapsed := time.Since(start)
			fmt.Printf("Функция выполнена за %s\n", elapsed)
			if err != nil {
				logger.Log.Debug("error func Run, method AddMessage by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			messageListByRoom, err := uc.db.SelectMessageListWithoutCtx(c.Message.RoomID)
			if err != nil {
				logger.Log.Debug(
					"error func Run, method SelectMessageListWithoutCtx by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			if _, ok := uc.hub.Clients[c.Message.RoomID]; ok {
				for _, cl := range uc.hub.Clients[c.Message.RoomID] {
					cl.Content <- &ws.Content{
						Message:           c.Message,
						MessageListByRoom: messageListByRoom,
					}
				}
			}
		}
	}
}

func (uc *UseCaseRoom) CreateRoom(ctx *fiber.Ctx, r CreateRoomRequest) (*ws.Room, error) {
	roomRequest := &ws.Room{
		RoomName: r.RoomName,
		Title:    r.Title,
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
	rp := ws.RoomProfile{
		RoomID:    newRoom.ID,
		ProfileID: profile.ID,
	}
	_, err = uc.db.AddRoomProfile(&rp)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method AddRoomProfile by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return newRoom, nil
}

func (uc *UseCaseRoom) GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error) {
	params := ws.QueryParamsRoomList{}
	if err := ctx.QueryParser(&params); err != nil {
		logger.Log.Debug("error func GetRoomList, method QueryParser by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	response, err := uc.db.SelectRoomList(ctx, &params)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return response, nil
}

func (uc *UseCaseRoom) GetRoomListByProfile(ctx *fiber.Ctx, r GetRoomListByProfileRequest) ([]*ws.RoomWithProfileResponse, error) {
	profileId, err := strconv.ParseInt(r.ProfileID, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetRoomListByProfile, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	response, err := uc.db.SelectRoomListByProfile(ctx, profileId)
	if err != nil {
		logger.Log.Debug("error func GetRoomListByProfile, method SelectList by path internal/useCase/room/room.go",
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
		Content:  make(chan *ws.Content),
	}
	profile, err := uc.db.FindProfile(userId)
	rp := ws.RoomProfile{
		RoomID:    roomId,
		ProfileID: profile.ID,
	}
	_, err = uc.db.AddRoomProfile(&rp)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method AddUser by path internal/useCase/room/room.go",
			zap.Error(err))
	}
	uc.hub.Register <- cl
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return ""
}

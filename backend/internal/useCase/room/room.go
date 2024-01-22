package room

import (
	"context"
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
	Limit  string `json:"limit"`
	Page   string `json:"page"`
}

func (uc *UseCaseRoom) Run(ctx context.Context) {
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
				IsJoined:  true,
				IsLeft:    false,
				Content:   cl.Username + " has joined the channel",
			}
			uc.hub.Broadcast <- &ws.Content{
				Message: message,
				Page:    cl.Page,
				Limit:   cl.Limit,
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
				IsJoined:  false,
				IsLeft:    true,
				Content:   cl.Username + " left the channel",
			}
			uc.hub.Broadcast <- &ws.Content{
				Message: message,
				Page:    cl.Page,
				Limit:   cl.Limit,
			}

		case c := <-uc.hub.Broadcast:
			fmt.Println("hub Broadcast: ", c)
			//start := time.Now()
			_, err := uc.db.AddMessage(ctx, c.Message)
			//elapsed := time.Since(start)
			//fmt.Printf("Функция выполнена за %s\n", elapsed)
			if err != nil {
				logger.Log.Debug("error func Run, method AddMessage by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			messageListByRoom, err := uc.db.SelectMessageList(
				ctx, c.Message.RoomID, c.Page, c.Limit)
			if err != nil {
				logger.Log.Debug(
					"error func Run, method SelectMessageList by path internal/useCase/room/room.go",
					zap.Error(err))
			}
			if _, ok := uc.hub.Clients[c.Message.RoomID]; ok {
				for _, cl := range uc.hub.Clients[c.Message.RoomID] {
					cl.Content <- &ws.Content{
						Message:           c.Message,
						MessageListByRoom: messageListByRoom,
						Page:              c.Page,
						Limit:             c.Limit,
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
	newRoom, err := uc.db.CreateRoom(ctx.Context(), roomRequest)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method Create by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	profile, err := uc.db.FindProfile(ctx.Context(), r.UserID)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method FindProfile by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	rp := ws.RoomProfile{
		RoomID:    newRoom.ID,
		ProfileID: profile.ID,
	}
	_, err = uc.db.AddRoomProfile(ctx.Context(), &rp)
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
	response, err := uc.db.SelectRoomList(ctx.Context(), &params)
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
	response, err := uc.db.SelectRoomListByProfile(ctx.Context(), profileId)
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
		logger.Log.Debug("error func GetUserList, method ParseInt by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	fmt.Println("roomId: ", roomId)
	clientList, err := uc.db.SelectUserList(ctx.Context())
	if err != nil {
		logger.Log.Debug("error func GetUserList, method SelectList by path internal/useCase/room/room.go",
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

func (uc *UseCaseRoom) GetMessageList(ctx *fiber.Ctx, r GetRoomMessagesRequest) (*ws.ResponseMessageList, error) {
	roomId, err := strconv.ParseInt(r.RoomID, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method ParseInt by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	page, err := strconv.ParseUint(r.Page, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method ParseUint by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	limit, err := strconv.ParseUint(r.Limit, 10, 64)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method ParseUint by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	messageList, err := uc.db.SelectMessageList(ctx.Context(), roomId, page, limit)
	if err != nil {
		logger.Log.Debug("error func GetMessageList, method SelectList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return messageList, nil
}

// checkAndAddCommonRoom - основная функция проверки и добавления записей
func (uc *UseCaseRoom) checkAndAddCommonRoom(
	ctx context.Context, r *ws.Room, senderId int64, receiverId int64) (*int64, error) {
	exists, roomID, err := uc.db.CheckIfCommonRoomExists(ctx, senderId, receiverId)
	if err != nil {
		return nil, err
	}
	fmt.Println("checkAndAddCommonRoom exists: ", exists)
	if !exists {
		roomCreated, err := uc.db.CreateRoom(ctx, r)
		if err != nil {
			return nil, err
		}
		roomID = roomCreated.ID
		err = uc.db.InsertRoomProfiles(ctx, roomID, senderId, receiverId)
		if err != nil {
			logger.Log.Debug(
				"error func checkAndAddCommonRoom, method AddUser by path internal/useCase/room/room.go",
				zap.Error(err))
		}
	}
	return &roomID, nil
}

func (uc *UseCaseRoom) JoinRoom(ctx context.Context, conn *websocket.Conn) string {
	userId := conn.Query("userId")
	username := conn.Query("username")
	roomTitle := conn.Query("roomTitle")
	receiverIdStr := conn.Query("receiverId")
	receiverId, err := strconv.ParseInt(receiverIdStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return ""
	}
	pageStr := conn.Query("page")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return ""
	}
	limitStr := conn.Query("limit")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method ParseInt roomIdStr by path internal/useCase/room/room.go",
			zap.Error(err))
		return ""
	}
	profile, err := uc.db.FindProfile(ctx, userId)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method FindProfile by path internal/useCase/room/room.go",
			zap.Error(err))
	}
	r := ws.Room{
		RoomName: username,
		Title:    roomTitle,
	}
	//checkAndAddCommonRoom - проверка и добавление записи
	roomID, err := uc.checkAndAddCommonRoom(ctx, &r, profile.ID, receiverId)
	if err != nil {
		logger.Log.Debug("error func JoinRoom, method checkAndAddCommonRoom by path internal/useCase/room/room.go",
			zap.Error(err))
	}
	cl := &ws.Client{
		RoomID:   *roomID,
		UserID:   userId,
		Username: username,
		Page:     page,
		Limit:    limit,
		Conn:     conn,
		Content:  make(chan *ws.Content),
	}
	uc.hub.Register <- cl
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return ""
}

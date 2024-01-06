package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

type JoinRoomUseCase struct {
	hub *ws.Hub
}

func NewJoinRoomUseCase(h *ws.Hub) *JoinRoomUseCase {
	return &JoinRoomUseCase{
		hub: h,
	}
}

func (uc *JoinRoomUseCase) JoinRoom(ctx *fiber.Ctx) (string, error) {
	upgrader := websocket.Upgrader{
		// ReadBufferSize - размер буфера чтения
		ReadBufferSize: 1024,
		// WriteBufferSize - размер буфера записи
		WriteBufferSize: 1024,
		// функция проверки источника, которая проверяет запрос и возвращает логическое значение
		CheckOrigin: func(r *http.Request) bool {
			// TODO: убрать return true, после тестирования в postman
			//origin := r.Header.Get("Origin")
			//return origin == "http://localhost:3000"
			return true
		},
	}
	// TODO: требуется получить conn
	conn, err := upgrader.Upgrade()
	if err != nil {
		return "", err
	}
	roomId := ctx.Params("roomId")
	params := ws.QueryParamsJoinRoom{}
	if err := ctx.QueryParser(&params); err != nil {
		logger.Log.Debug("error in method ctx.QueryParser", zap.Error(err))
		return "", err
	}
	cl := &ws.Client{
		ID:       params.ClientID,
		RoomID:   roomId,
		Username: params.Username,
		Conn:     conn,
		Message:  make(chan *ws.Message),
	}
	m := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomId,
		Username: params.Username,
	}
	// Register a new client through the register channel
	uc.hub.Register <- cl
	// Broadcast that message
	uc.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return "is joined", nil
}

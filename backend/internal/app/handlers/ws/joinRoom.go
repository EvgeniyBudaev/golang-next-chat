package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/contrib/websocket"
)

type IJoinRoomUseCase interface {
	JoinRoom(conn *websocket.Conn) string
}

func JoinRoomHandler(uc IJoinRoomUseCase) func(c *websocket.Conn) {
	return func(conn *websocket.Conn) {
		logger.Log.Info("GET /api/v1/ws/room/join/:roomId?username=user")
		uc.JoinRoom(conn)
	}
}

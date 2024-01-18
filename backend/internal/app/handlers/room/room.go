package room

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	roomUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/room"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type UseCaseRoom interface {
	CreateRoom(ctx *fiber.Ctx, r roomUseCase.CreateRoomRequest) (*ws.Room, error)
	GetUserList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error)
	GetMessageList(ctx *fiber.Ctx, r roomUseCase.GetRoomMessagesRequest) ([]*ws.ResponseMessage, error)
	GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error)
	JoinRoom(conn *websocket.Conn) string
}

func CreateRoomHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("POST /api/v1/room/create")
		req := roomUseCase.CreateRoomRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			logger.Log.Debug("error func CreateRoomHandler,"+
				" method ctx.BodyParse by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := uc.CreateRoom(ctx, req)
		if err != nil {
			logger.Log.Debug("error func CreateRoomHandler,"+
				" method uc.CreateRoom by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctx, response)
	}
}

func GetUserListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/:roomId/client/list")
		response, err := uc.GetUserList(ctx)
		if err != nil {
			logger.Log.Debug(
				"error func GetUserListHandler,"+
					" method uc.GetUserList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

func GetMessageListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/message/list")
		var request = roomUseCase.GetRoomMessagesRequest{}
		err := ctx.BodyParser(&request)
		if err != nil {
			logger.Log.Debug(
				"error func GetMessageListHandler,"+
					" method BodyParser by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := uc.GetMessageList(ctx, request)
		if err != nil {
			logger.Log.Debug(
				"error func GetUserListHandler,"+
					" method uc.GetUserList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

func GetRoomListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/list")
		response, err := uc.GetRoomList(ctx)
		if err != nil {
			logger.Log.Debug("error func GetRoomListHandler,"+
				" method uc.GetRoomList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

func JoinRoomHandler(uc UseCaseRoom) func(c *websocket.Conn) {
	return func(conn *websocket.Conn) {
		logger.Log.Info("GET /api/v1/room/join/:roomId?userId=userId&username=username")
		uc.JoinRoom(conn)
	}
}

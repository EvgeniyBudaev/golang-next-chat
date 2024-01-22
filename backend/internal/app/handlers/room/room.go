package room

import (
	"context"
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
	CreateRoom(ctf *fiber.Ctx, r roomUseCase.CreateRoomRequest) (*ws.Room, error)
	GetUserList(ctf *fiber.Ctx) ([]*ws.ClientResponse, error)
	GetMessageList(ctf *fiber.Ctx, r roomUseCase.GetRoomMessagesRequest) (*ws.ResponseMessageList, error)
	GetRoomList(ctf *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error)
	GetRoomListByProfile(
		ctf *fiber.Ctx, r roomUseCase.GetRoomListByProfileRequest) ([]*ws.RoomWithProfileResponse, error)
	JoinRoom(ctx context.Context, conn *websocket.Conn) string
}

func CreateRoomHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("POST /api/v1/room/create")
		req := roomUseCase.CreateRoomRequest{}
		if err := ctf.BodyParser(&req); err != nil {
			logger.Log.Debug("error func CreateRoomHandler,"+
				" method ctx.BodyParse by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		response, err := uc.CreateRoom(ctf, req)
		if err != nil {
			logger.Log.Debug("error func CreateRoomHandler,"+
				" method uc.CreateRoom by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctf, response)
	}
}

func GetUserListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/:roomId/client/list")
		response, err := uc.GetUserList(ctf)
		if err != nil {
			logger.Log.Debug(
				"error func GetUserListHandler,"+
					" method uc.GetUserList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctf, response)
	}
}

func GetMessageListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/message/list")
		var request = roomUseCase.GetRoomMessagesRequest{}
		err := ctf.BodyParser(&request)
		if err != nil {
			logger.Log.Debug(
				"error func GetMessageListHandler,"+
					" method BodyParser by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		response, err := uc.GetMessageList(ctf, request)
		if err != nil {
			logger.Log.Debug(
				"error func GetUserListHandler,"+
					" method uc.GetUserList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctf, response)
	}
}

func GetRoomListHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/list")
		response, err := uc.GetRoomList(ctf)
		if err != nil {
			logger.Log.Debug("error func GetRoomListHandler,"+
				" method uc.GetRoomList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctf, response)
	}
}

func GetRoomListByProfileHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/profile/room/list")
		req := roomUseCase.GetRoomListByProfileRequest{}
		if err := ctf.BodyParser(&req); err != nil {
			logger.Log.Debug("error func GetRoomListByProfileHandler,"+
				" method ctx.BodyParse by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		response, err := uc.GetRoomListByProfile(ctf, req)
		if err != nil {
			logger.Log.Debug("error func GetRoomListByProfileHandler,"+
				" method uc.GetRoomList by path internal/handlers/room/room.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctf, response)
	}
}

func JoinRoomHandler(uc UseCaseRoom) fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		ctx := ctf.Context()
		return websocket.New(func(conn *websocket.Conn) {
			logger.Log.Info("WS /api/v1/room/join?userId=?&username=?&roomTitle=?&receiverId=?")
			uc.JoinRoom(ctx, conn)
		})(ctf)
	}
}

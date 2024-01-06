package ws

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type IGetRoomListUseCase interface {
	GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomResponse, error)
}

func GetRoomListHandler(uc IGetRoomListUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/ws/room/list")
		response, err := uc.GetRoomList(ctx)
		if err != nil {
			logger.Log.Debug("error in method uc.GetRoomList", zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

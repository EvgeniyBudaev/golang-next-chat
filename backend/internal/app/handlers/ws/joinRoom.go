package ws

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type IJoinRoomUseCase interface {
	JoinRoom(ctx *fiber.Ctx) (string, error)
}

func JoinRoomHandler(uc IJoinRoomUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/ws/room/join/:roomId?username=user")
		response, err := uc.JoinRoom(ctx)
		if err != nil {
			logger.Log.Debug("error in method uc.JoinRoom", zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

package ws

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	wsUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/ws"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type ICreateRoomUseCase interface {
	CreateRoom(ctx *fiber.Ctx, r wsUseCase.CreateRoomRequest) (string, error)
}

func CreateRoomHandler(uc ICreateRoomUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("POST /api/v1/ws/room/create")
		req := wsUseCase.CreateRoomRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			logger.Log.Debug("error in method ctx.BodyParse", zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := uc.CreateRoom(ctx, req)
		if err != nil {
			logger.Log.Debug("error in method uc.CreateRoom", zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctx, response)
	}
}

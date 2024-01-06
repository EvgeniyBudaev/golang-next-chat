package ws

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type IGetClientListUseCase interface {
	GetClientList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error)
}

func GetClientListHandler(uc IGetClientListUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/room/:roomId/client/list")
		response, err := uc.GetClientList(ctx)
		if err != nil {
			logger.Log.Debug("error in method uc.GetClientList", zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctx, response)
	}
}

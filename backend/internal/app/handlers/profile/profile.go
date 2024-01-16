package profile

import (
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	profileUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/profile"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type HandlerProfile struct {
	uc *profileUseCase.UseCaseProfile
}

func NewHandlerProfile(uc *profileUseCase.UseCaseProfile) *HandlerProfile {
	return &HandlerProfile{uc: uc}
}

func (h *HandlerProfile) CreateProfileHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("POST /api/v1/profile/create")
		req := profileUseCase.CreateProfileRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			logger.Log.Debug(
				"error func CreateProfileHandler,"+
					" method ctx.BodyParse by path internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := h.uc.CreateProfile(ctx, req)
		if err != nil {
			logger.Log.Debug(
				"error func CreateProfileHandler, method uc.CreateRoom by path"+
					" internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctx, response)
	}
}

func (h *HandlerProfile) GetProfileByUUIDHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/profile/detail")
		req := profileUseCase.GetProfileRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			logger.Log.Debug(
				"error func GetProfileByUUIDHandler,"+
					" method ctx.BodyParse by path internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := h.uc.GetProfileByUUID(ctx, req)
		if err != nil {
			logger.Log.Debug(
				"error func GetProfileByUUIDHandler, method GetProfileByUUID by path"+
					" internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctx, response)
	}
}

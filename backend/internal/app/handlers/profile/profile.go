package profile

import (
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
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

func (h *HandlerProfile) GetProfileByUsernameHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/profile/detail")
		req := profileUseCase.GetProfileRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			logger.Log.Debug(
				"error func GetProfileByUsernameHandler,"+
					" method ctx.BodyParse by path internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		response, err := h.uc.GetProfileByUsername(ctx, req)
		if err != nil {
			logger.Log.Debug(
				"error func GetProfileByUsernameHandler, method GetProfileByUsername by path"+
					" internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctx, err, http.StatusBadRequest)
		}
		return r.WrapCreated(ctx, response)
	}
}

func (h *HandlerProfile) GetProfileListHandler() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		logger.Log.Info("GET /api/v1/profile/list")
		req := profileEntity.QueryParamsProfileList{}
		if err := ctf.BodyParser(&req); err != nil {
			logger.Log.Debug(
				"error func GetProfileListHandler,"+
					" method ctx.BodyParse by path internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		response, err := h.uc.GetProfileList(ctf, &req)
		if err != nil {
			logger.Log.Debug(
				"error func GetProfileListHandler, method GetProfileList by path"+
					" internal/handlers/profile/profile.go",
				zap.Error(err))
			return r.WrapError(ctf, err, http.StatusBadRequest)
		}
		return r.WrapOk(ctf, response)
	}
}

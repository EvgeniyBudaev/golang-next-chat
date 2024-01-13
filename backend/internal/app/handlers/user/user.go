package user

import (
	"context"
	r "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/response"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type IUserUseCase interface {
	Register(ctx context.Context, request user.RegisterRequest) (*user.RegisterResponse, error)
	GetUserList(ctx context.Context, query user.QueryParamsUserList) ([]*gocloak.User, error)
}

func PostRegisterHandler(uc IUserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		logger.Log.Info("POST /api/v1/user/user")
		var request = user.RegisterRequest{}
		err := c.BodyParser(&request)
		if err != nil {
			logger.Log.Debug("error while PostRegisterHandler. Error in BodyParser", zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		response, err := uc.Register(ctx, request)
		if err != nil {
			logger.Log.Debug("error while PostRegisterHandler. Error in Register", zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		return r.WrapCreated(c, response)
	}
}

func GetUserListHandler(uc IUserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		logger.Log.Info("GET /api/v1/user/list")
		query := user.QueryParamsUserList{}
		if err := c.QueryParser(&query); err != nil {
			logger.Log.Debug("error while GetCatalogList. error in method QueryParser", zap.Error(err))
			return err
		}
		response, err := uc.GetUserList(ctx, query)
		if err != nil {
			logger.Log.Debug("error while PostRegisterHandler. Error in Register", zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		return r.WrapOk(c, response)
	}
}

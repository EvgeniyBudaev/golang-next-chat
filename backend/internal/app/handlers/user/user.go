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

type UseCaseUser interface {
	Register(ctx context.Context, request user.RegisterRequest) (*gocloak.User, error)
	GetUserList(ctx context.Context, query user.QueryParamsUserList) ([]*gocloak.User, error)
}

func PostRegisterHandler(uc UseCaseUser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		logger.Log.Info("POST /api/v1/user/user")
		var request = user.RegisterRequest{}
		err := c.BodyParser(&request)
		if err != nil {
			logger.Log.Debug("error func PostRegisterHandler, method BodyParser by path handlers/user/user.go",
				zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		response, err := uc.Register(ctx, request)
		if err != nil {
			logger.Log.Debug("error func PostRegisterHandler, method Register by path handlers/user/user.go",
				zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		return r.WrapCreated(c, response)
	}
}

func GetUserListHandler(uc UseCaseUser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		logger.Log.Info("GET /api/v1/user/list")
		query := user.QueryParamsUserList{}
		if err := c.QueryParser(&query); err != nil {
			logger.Log.Debug("error func GetUserListHandler, method QueryParser", zap.Error(err))
			return err
		}
		response, err := uc.GetUserList(ctx, query)
		if err != nil {
			logger.Log.Debug("error func GetUserListHandler, method Register", zap.Error(err))
			return r.WrapError(c, err, http.StatusBadRequest)
		}
		return r.WrapOk(c, response)
	}
}

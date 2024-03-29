package user

import (
	"context"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/searching"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
	"strings"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobileNumber"`
}

type QueryParamsUserList struct {
	searching.Searching
}

type UseCaseUser struct {
	identity Identity
}

func NewUserUseCase(i Identity) *UseCaseUser {
	return &UseCaseUser{
		identity: i,
	}
}

func (uc *UseCaseUser) Register(ctx context.Context, request RegisterRequest) (*gocloak.User, error) {
	var user = gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}
	if strings.TrimSpace(request.MobileNumber) != "" {
		(*user.Attributes)["mobileNumber"] = []string{request.MobileNumber}
	}
	response, err := uc.identity.CreateUser(ctx, user, request.Password, "customer")
	if err != nil {
		logger.Log.Debug("error in method CreateUser by path useCase/user/user.go", zap.Error(err))
		return nil, err
	}
	return response, nil
}

func (uc *UseCaseUser) GetUserList(ctx context.Context, query QueryParamsUserList) ([]*gocloak.User, error) {
	response, err := uc.identity.GetUserList(ctx, query)
	if err != nil {
		logger.Log.Debug("error in method GetUserList by path useCase/user/user.go", zap.Error(err))
		return nil, err
	}
	return response, nil
}

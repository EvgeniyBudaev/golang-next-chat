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

type RegisterResponse struct {
	User *gocloak.User
}

type QueryParamsUserList struct {
	searching.Searching
}

type UserListResponse struct {
	UserList []*gocloak.User `json:"userList"`
}

type UserUseCase struct {
	identity IIdentity
}

func NewUserUseCase(i IIdentity) *UserUseCase {
	return &UserUseCase{
		identity: i,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {
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
	userResponse, err := uc.identity.CreateUser(ctx, user, request.Password, "customer")
	if err != nil {
		logger.Log.Debug("error in method CreateUser by path useCase/user/user.go", zap.Error(err))
		return nil, err
	}
	var response = &RegisterResponse{User: userResponse}
	return response, nil
}

func (uc *UserUseCase) GetUserList(ctx context.Context, query QueryParamsUserList) ([]*gocloak.User, error) {
	response, err := uc.identity.GetUserList(ctx, query)
	if err != nil {
		logger.Log.Debug("error in method GetUserList by path useCase/user/user.go", zap.Error(err))
		return nil, err
	}
	return response, nil
}

package routes

import (
	registerHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/register"
	wsHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	identityEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	wsEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	userUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
	wsUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/ws"
	"github.com/gofiber/fiber/v2"
)

var (
	prefix = "/api/v1"
)

func InitPublicRoutes(app *fiber.App, config *config.Config) {
	app.Static("/static", "./static")
	// store
	identityManager := identityEntity.NewIdentity(config)
	// hub
	hub := wsEntity.NewHub()
	go hub.Run()
	// useCase
	useCaseRegister := userUseCase.NewRegisterUseCase(identityManager)
	useCaseCreateRoom := wsUseCase.NewCreateRoomUseCase(hub)
	useCaseJoinRoom := wsUseCase.NewJoinRoomUseCase(hub)
	useCaseGetRoomList := wsUseCase.NewGetRoomListUseCase(hub)
	useCaseGetClientList := wsUseCase.NewGetClientListUseCase(hub)
	// handlers
	grp := app.Group(prefix)
	grp.Post("/user/register", registerHandler.PostRegisterHandler(useCaseRegister))
	grp.Post("/ws/room/create", wsHandler.CreateRoomHandler(useCaseCreateRoom))
	grp.Get("/ws/room/join/:roomId", wsHandler.JoinRoomHandler(useCaseJoinRoom))
	grp.Get("/ws/room/list", wsHandler.GetRoomListHandler(useCaseGetRoomList))
	grp.Get("/ws/room/:roomId/client/list", wsHandler.GetClientListHandler(useCaseGetClientList))
}

func InitProtectedRoutes(app *fiber.App, config *config.Config) {
}

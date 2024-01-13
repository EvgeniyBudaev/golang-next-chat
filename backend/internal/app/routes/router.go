package routes

import (
	userHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/user"
	wsHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	identityEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	wsEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	userUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
	wsUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/ws"
	"github.com/gofiber/contrib/websocket"
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
	useCaseUser := userUseCase.NewUserUseCase(identityManager)
	useCaseCreateRoom := wsUseCase.NewCreateRoomUseCase(hub)
	useCaseJoinRoom := wsUseCase.NewJoinRoomUseCase(hub)
	useCaseGetRoomList := wsUseCase.NewGetRoomListUseCase(hub)
	useCaseGetClientList := wsUseCase.NewGetClientListUseCase(hub)
	// handlers
	grp := app.Group(prefix)

	app.Use("/ws/room/join/:roomId", func(ctx *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}

		return ctx.Next()
	})

	grp.Post("/user/register", userHandler.PostRegisterHandler(useCaseUser))
	grp.Get("/user/list", userHandler.GetUserListHandler(useCaseUser))
	grp.Post("/ws/room/create", wsHandler.CreateRoomHandler(useCaseCreateRoom))
	grp.Get("/ws/room/join/:roomId", websocket.New(wsHandler.JoinRoomHandler(useCaseJoinRoom)))
	grp.Get("/ws/room/list", wsHandler.GetRoomListHandler(useCaseGetRoomList))
	grp.Get("/ws/room/:roomId/client/list", wsHandler.GetClientListHandler(useCaseGetClientList))
}

func InitProtectedRoutes(app *fiber.App, config *config.Config) {
}

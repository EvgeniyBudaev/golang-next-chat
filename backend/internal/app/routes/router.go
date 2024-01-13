package routes

import (
	wsHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/room"
	userHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/user"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	identityEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	wsEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	wsUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/room"
	userUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
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
	useCaseRoom := wsUseCase.NewUseCaseRoom(hub)
	// handlers
	grp := app.Group(prefix)

	app.Use("/room/room/join/:roomId", func(ctx *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}

		return ctx.Next()
	})

	grp.Post("/user/register", userHandler.PostRegisterHandler(useCaseUser))
	grp.Get("/user/list", userHandler.GetUserListHandler(useCaseUser))
	grp.Post("/room/create", wsHandler.CreateRoomHandler(useCaseRoom))
	grp.Get("/room/join/:roomId", websocket.New(wsHandler.JoinRoomHandler(useCaseRoom)))
	grp.Get("/room/list", wsHandler.GetRoomListHandler(useCaseRoom))
	grp.Get("/room/:roomId/client/list", wsHandler.GetClientListHandler(useCaseRoom))
}

func InitProtectedRoutes(app *fiber.App, config *config.Config) {
}

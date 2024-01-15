package routes

import (
	profileHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/profile"
	roomHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/room"
	userHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/user"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/room"
	identityEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	wsEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	profileUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/profile"
	wsUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/room"
	userUseCase "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var (
	prefix = "/api/v1"
)

func InitPublicRoutes(app *fiber.App, config *config.Config, db *db.Database) {
	app.Static("/static", "./static")
	// store
	identityManager := identityEntity.NewIdentity(config)
	// db
	dbProfile := profile.NewPGProfileDB(db.GetDB())
	dbRoom := room.NewPGRoomDB(db.GetDB())
	// hub
	hub := wsEntity.NewHub()
	// useCase
	useCaseUser := userUseCase.NewUserUseCase(identityManager)
	useCaseRoom := wsUseCase.NewUseCaseRoom(hub, dbRoom)
	app.Use(func(ctx *fiber.Ctx) error {
		go useCaseRoom.Run(ctx)
		return ctx.Next()
	})
	useCaseProfile := profileUseCase.NewUseCaseProfile(dbProfile)
	// handlers
	grp := app.Group(prefix)
	ph := profileHandler.NewHandlerProfile(useCaseProfile)

	app.Use("/room/join/:roomId", func(ctx *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}
		return ctx.Next()
	})

	grp.Post("/user/register", userHandler.PostRegisterHandler(useCaseUser))
	grp.Get("/user/list", userHandler.GetUserListHandler(useCaseUser))
	grp.Post("/room/create", roomHandler.CreateRoomHandler(useCaseRoom))
	grp.Get("/room/join/:roomId", websocket.New(roomHandler.JoinRoomHandler(useCaseRoom)))
	grp.Get("/room/list", roomHandler.GetRoomListHandler(useCaseRoom))
	grp.Get("/room/:roomId/user/list", roomHandler.GetUserListHandler(useCaseRoom))
	grp.Post("/room/message/list", roomHandler.GetMessageListHandler(useCaseRoom))
	grp.Post("/profile/create", ph.CreateProfileHandler())
	grp.Get("/profile/uuid/:uuid", ph.GetProfileByUUIDHandler())
}

func InitProtectedRoutes(app *fiber.App, config *config.Config, db *db.Database) {
	// app.Use("/static", middlewares.NewRequiresRealmRole("admin"), filesystem.New(filesystem.Config{
	// 	Root: http.Dir("./static"),
	// }))
}

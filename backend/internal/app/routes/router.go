package routes

import (
	registerHandler "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/handlers/register"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/useCase/user"
	"github.com/gofiber/fiber/v2"
)

var (
	prefix string = "/api/v1"
)

func InitPublicRoutes(app *fiber.App, config *config.Config) {
	app.Static("/static", "./static")
	grp := app.Group(prefix)

	// store
	identityManager := identity.NewIdentity(config)

	// useCase
	useCaseRegister := user.NewRegisterUseCase(identityManager)

	// handlers
	grp.Post("/user/register", registerHandler.PostRegisterHandler(useCaseRegister))
}

func InitProtectedRoutes(app *fiber.App, config *config.Config) {
}

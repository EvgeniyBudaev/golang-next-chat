package middlewares

import (
	"context"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/identity"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/shared/enums"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitFiberMiddlewares(app *fiber.App,
	cfg *config.Config,
	initPublicRoutes func(app *fiber.App, config *config.Config),
	initProtectedRoutes func(app *fiber.App, config *config.Config)) {
	app.Use(requestid.New())
	app.Use(func(c *fiber.Ctx) error {
		// get the request id that was added by requestid middleware
		var requestId = c.Locals("requestid")
		// create a new context and add the requestid to it
		var ctx = context.WithValue(context.Background(), enums.ContextKeyRequestId, requestId)
		c.SetUserContext(ctx)
		return c.Next()
	})
	// routes that don't require a JWT token
	initPublicRoutes(app, cfg)
	tokenRetrospector := identity.NewIdentity(cfg)
	app.Use(NewJwtMiddleware(cfg, tokenRetrospector))
	// routes that require authentication/authorization
	initProtectedRoutes(app, cfg)
}

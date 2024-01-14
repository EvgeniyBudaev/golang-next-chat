package app

import (
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app/routes"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/config"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func Start() error {
	// Initialization
	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,
	})
	// Config
	cfg, err := config.Load()
	if err != nil {
		logger.Log.Debug("error func Start, method Load by path internal/app/app.go", zap.Error(err))
		return err
	}
	// Logging
	if err := logger.Initialize(cfg.LoggerLevel); err != nil {
		logger.Log.Debug("error func Start, method Initialize by path internal/app/app.go", zap.Error(err))
		return err
	}
	// Database
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
		cfg.DBSSlMode)
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		logger.Log.Debug("error func Start, method Open by path internal/app/app.go", zap.Error(err))
		return err
	}
	database := db.NewDatabase(conn)
	err = conn.Ping()
	if err != nil {
		logger.Log.Debug("error func Start, method Ping by path internal/app/app.go", zap.Error(err))
		return err
	}
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, X-Requested-With, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	// Routes
	middlewares.InitFiberMiddlewares(app, cfg, database, routes.InitPublicRoutes, routes.InitProtectedRoutes)
	return app.Listen(cfg.Port)
}

package main

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/app"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Log.Debug("error in method app.Start", zap.Error(err))
		log.Fatal(err)
	}
}

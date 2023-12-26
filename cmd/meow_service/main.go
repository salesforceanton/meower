package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/salesforceanton/meower/internal/config"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/meow_service"
	"github.com/salesforceanton/meower/internal/repository"
)

func main() {
	// Configure block
	logger.ConfigureLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with config initialization")
		return
	}
	logger.LogInfo("Configs initialized")

	// Connect to Database
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Database connect")
		return
	}
	logger.LogInfo("Connected to DB")

	// Connect to Event Bus
	eventbus, err := eventbus.NewNatsEventbus(fmt.Sprintf("%s:%s", cfg.NatsHost, cfg.NatsPort))
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Event Bus connect")
		return
	}
	logger.LogInfo("Connected to Event Bus")

	// Init Deps
	repo := repository.NewPostgresRepo(db)
	handler := meow_service.NewHandler(repo, eventbus)

	// Run server
	server := new(meow_service.MeowServiceServer)
	go func() {
		if err = server.Run(cfg.Port, handler.InitRoutes()); err != nil {
			logger.LogError(err.Error(), "[Runtime Error]: Error with Meow Server Running or Server has been Stopped")
			return
		}
	}()

	// Shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	server.Shutdown(context.Background())
	repo.Close()
	eventbus.Close()

	logger.LogInfo("Server shutdown successfully")
}

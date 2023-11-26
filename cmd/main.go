package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/salesforceanton/meower/internal/config"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/repository"
)

func main() {
	// Configure block
	logger.ConfigureLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogError("Error with config initialization", err.Error())
		return
	}
	logger.LogInfo("Configs initialized")

	// Database block
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		fmt.Println(err)
		logger.LogError("Error with database connect", err.Error())
		return
	}

	repo := repository.NewPostgresRepo(db)
	logger.LogInfo("Connected to DB")

	// Shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	repo.Close()

	logger.LogInfo("Server shutdown successfully")
}

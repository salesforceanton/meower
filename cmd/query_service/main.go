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
	"github.com/salesforceanton/meower/internal/query_service"
	"github.com/salesforceanton/meower/internal/repository"
	"github.com/salesforceanton/meower/internal/schema"
	search_repo "github.com/salesforceanton/meower/internal/search"
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

	// Connect to Search Repo
	searchRepo, err := search_repo.NewElasticRepo(fmt.Sprintf("%s:%s", cfg.ElasticsearchHost, cfg.ElasticsearchPort))
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Search Repo connect")
		return
	}
	logger.LogInfo("Connected to Search Repo")

	// Connect to Event Bus
	mq, err := eventbus.NewNatsEventbus(fmt.Sprintf("%s:%s", cfg.NatsHost, cfg.NatsPort))
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Event Bus connect")
		return
	}
	logger.LogInfo("Connected to Event Bus")

	// Subscribe Meow Created
	err = mq.SubscribeMeowCreated(func(m eventbus.MeowCreatedMessage) {
		message := schema.Meow{
			Id:        m.Id,
			Body:      m.Body,
			CreatedAt: m.CreatedAt,
		}
		if err := searchRepo.InsertMeow(context.TODO(), message); err != nil {
			logger.LogError(err.Error(), "[Event Bus Subscription]: error onmeow create event handler")
		}
	})
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Event Bus Subscribe")
		return
	}
	logger.LogInfo("Subscribe on meowcreated")

	// Init Deps
	repo := repository.NewPostgresRepo(db)
	handler := query_service.NewHandler(repo, mq, searchRepo)

	// Run server
	server := new(query_service.QueryServiceServer)
	go func() {
		if err = server.Run(cfg.Port, handler.InitRoutes()); err != nil {
			logger.LogError(err.Error(), "[Runtime Error]: Error with Query Server Running or Server has been Stopped")
			return
		}
	}()

	// Shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	server.Shutdown(context.Background())
	repo.Close()
	mq.Close()

	logger.LogInfo("Server shutdown successfully")
}

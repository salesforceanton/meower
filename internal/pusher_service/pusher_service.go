package pusher_service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/salesforceanton/meower/internal/config"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/web_socket"
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

	// Connect to Event Bus
	hub := web_socket.NewHub()
	mq, err := eventbus.NewNatsEventbus(fmt.Sprintf("%s:%s", cfg.NatsHost, cfg.NatsPort))
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Event Bus connect")
		return
	}
	logger.LogInfo("Connected to Event Bus")

	// Subscribe Meow Created
	err = mq.SubscribeMeowCreated(func(m eventbus.MeowCreatedMessage) {
		message := web_socket.NewMeowCreatedMessage(m.Id, m.Body, m.CreatedAt)

		hub.Broadcast(message)
	})
	if err != nil {
		logger.LogError(err.Error(), "[Runtime Error]: Error with Event Bus Subscribe")
		return
	}
	logger.LogInfo("Subscribe on meowcreated")

	// Init Deps
	handler := NewHandler(mq, hub)
	server := new(PusherServiceServer)

	// Run web socket Hub
	go hub.Run()

	// Run server
	go func() {
		if err = server.Run(cfg.Port, handler.InitRoutes()); err != nil {
			logger.LogError(err.Error(), "[Runtime Error]: Error with Pusher Server Running or Server has been Stopped")
			return
		}
	}()

	// Shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	server.Shutdown(context.Background())
	mq.Close()

	logger.LogInfo("Server shutdown successfully")
}

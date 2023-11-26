package main

import (
	"fmt"

	"github.com/salesforceanton/meower/internal/config"
	"github.com/salesforceanton/meower/internal/logger"
)

func main() {
	logger.ConfigureLogger()

	config, err := config.InitConfig()
	if err != nil {
		logger.LogError("Error with config initialization", err.Error())
	}

	fmt.Println(config)
}

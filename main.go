package main

import (
	"fmt"

	"github.com/wander2583/api-oportunits/config"
	"github.com/wander2583/api-oportunits/router"
)

var (
	logger config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		panic(err)
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}

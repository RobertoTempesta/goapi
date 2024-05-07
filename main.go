package main

import (
	"github.com/RobertoTempesta/goapi/config"
	"github.com/RobertoTempesta/goapi/router"
)

var (
	logger config.Logger
)

func main() {
	//Initialize logger
	logger = *config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	// Initialize Router
	// Obs: export é pela inicial do function sempre maiuscula
	router.Initialize()
}

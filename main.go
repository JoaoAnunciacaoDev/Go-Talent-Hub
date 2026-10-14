package main

import (
	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/config"
	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err == nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}

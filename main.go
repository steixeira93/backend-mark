package main

import (
	"github.com/steixeira93/mark/config"
	"github.com/steixeira93/mark/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialize()
}

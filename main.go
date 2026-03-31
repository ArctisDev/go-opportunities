package main

import (
	"github.com/ArctisDev/go-opportunities/config"
	"github.com/ArctisDev/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize configs
	err := config.InitDB()
	if err != nil {
		logger.Errorf("Failed config initialization: %v", err)
		return
	}

	// Initialize the router
	router.InitializeRouter()
}

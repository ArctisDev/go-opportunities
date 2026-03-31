package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {

	//  InitializeRouter sets up the Gin router.
	router := gin.Default()

	// InitializeRoutes defines the API routes.
	InitializeRoutes(router)
	
	// Start the server on port 8080.
	router.Run((":8080"))
}

package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
		router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.Run((":8080"))
}
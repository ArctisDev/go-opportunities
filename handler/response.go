package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Contet-type", "aplication/json")
	ctx.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Contet-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"messsage": fmt.Sprintf("operation from handler: %s sucessfully executed", op),
		"data":    data,
	})
}

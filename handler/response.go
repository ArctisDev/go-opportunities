package handler

import (
	"fmt"
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
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
		"data":     data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                  `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
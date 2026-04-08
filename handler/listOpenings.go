package handler

import (
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error listening openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error listening openings")
	}

	sendSucess(ctx, "list-openings", openings)
}

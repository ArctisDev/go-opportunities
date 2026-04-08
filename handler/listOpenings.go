package handler

import (
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary List Openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error listening openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error listening openings")
	}

	sendSucess(ctx, "list-openings", openings)
}

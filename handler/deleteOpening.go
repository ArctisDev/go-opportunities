package handler

import (
	"fmt"
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Delete Opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("Missing required query parameter: id")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Opening with id %s not found", id)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("Error deleting opening with id %s: %v", id, err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id: %s", id))
	}

	sendSucess(ctx, "delete-opening", opening)
}

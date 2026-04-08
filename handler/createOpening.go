package handler

import (
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:        request.Role,
		Company:     request.Company,
		Description: request.Description,
		Location:    request.Location,
		URL:         request.URL,
		Salary:      request.Salary,
		Remote:      *request.Remote,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	sendSucess(ctx, "CreateOpeningHandler", opening)
}

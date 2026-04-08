package handler

import (
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

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

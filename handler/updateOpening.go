package handler

import (
	"net/http"

	"github.com/ArctisDev/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Update Opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Failed to validate request body: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}
	// Update fields if they are provided in the request
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Description != "" {
		opening.Description = request.Description
	}
	if request.URL != "" {
		opening.URL = request.URL
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save Opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Failed to update opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Failed to update opening")
		return
	}

	sendSucess(ctx, "update-opening", opening)

}

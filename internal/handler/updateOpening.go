package handler

import (
	"fmt"
	"net/http"

	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		sendError(
			context,
			http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error(),
		)

		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			context,
			http.StatusNotFound,
			fmt.Sprintf("opening with id %s not found", id),
		)
		return
	}

	if request.Title != "" {
		opening.Title = request.Title
	}

	if request.Description != "" {
		opening.Description = request.Description
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.IsRemote != nil {
		opening.IsRemote = *request.IsRemote
	}

	if request.IsActive != nil {
		opening.IsActive = *request.IsActive
	}

	if request.IsSecret != nil {
		opening.IsSecret = *request.IsSecret
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(context, "updateOpening", opening)
}

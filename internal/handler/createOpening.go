package handler

import (
	"net/http"

	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Title: request.Title,
		Description: request.Description,
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		IsRemote: *request.IsRemote,
		IsActive: *request.IsActive,
		IsSecret: *request.IsSecret,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(context, "createOpening", opening)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context) {
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
		sendError(context,
			http.StatusNotFound,
			fmt.Sprintf("opening with id %s not found", id),
		)

		return
	}

	sendSuccess(context, "showOpening", opening)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(
			context, 
			http.StatusInternalServerError, 
			(fmt.Sprintf("error listing openings: %v", err.Error())),
		)
		return
	}

	sendSuccess(context, "listOpenings", openings)
}

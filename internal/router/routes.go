package router

import (
	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{

		v1.GET("/opening", func(context *gin.Context) { handler.ShowOpeningHandler(context) })
		v1.GET("/openings", func(context *gin.Context) { handler.ListOpeningsHandler(context) })
		v1.POST("/opening", func(context *gin.Context) { handler.CreateOpeningHandler(context) })
		v1.PUT("/opening", func(context *gin.Context) { handler.UpdateOpeningHandler(context) })
		v1.DELETE("/opening", func(context *gin.Context) { handler.DeleteOpeningHandler(context) })

	}
}

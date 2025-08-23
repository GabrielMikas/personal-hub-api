package router

import (
	"github.com/gin-gonic/gin"

	"github.com/GabrielMikas/personal-hub-api/handlers"
	cards_handler "github.com/GabrielMikas/personal-hub-api/handlers/cards"
	investments_handler "github.com/GabrielMikas/personal-hub-api/handlers/investments"
	savings_handler "github.com/GabrielMikas/personal-hub-api/handlers/savings"
)

func serveRoutes(router *gin.Engine) {
	handlers.InitializeHandlers()
	investments := router.Group("/finance/investments")
	{
		investments.GET("/", investments_handler.GetHandler)
		investments.POST("/", investments_handler.PostHandler)

	}
	savings := router.Group("/finance/savings")
	{
		savings.GET("/", savings_handler.GetHandler)
		savings.POST("/", savings_handler.PostHandler)
	}
	cards := router.Group("/hobbies/cards")
	{
		cards.GET("/", cards_handler.GetHandler)
		cards.POST("/", cards_handler.PostHandler)
	}
}

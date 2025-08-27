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
		investments.GET("/:id", investments_handler.GetById)
		investments.POST("/", investments_handler.PostHandler)
		investments.DELETE("/:id", investments_handler.DeleteHandler)
		investments.PATCH("/:id", investments_handler.PatchHandler)

	}
	savings := router.Group("/finance/savings")
	{
		savings.GET("/", savings_handler.GetHandler)
		savings.GET("/:id", savings_handler.GetById)
		savings.POST("/", savings_handler.PostHandler)
		savings.DELETE("/:id", savings_handler.DeleteHandler)
		savings.PATCH("/:id", savings_handler.PatchHandler)
	}
	cards := router.Group("/hobbies/cards")
	{
		cards.GET("/", cards_handler.GetHandler)
		cards.GET("/:id", cards_handler.GetById)
		cards.DELETE("/:id", cards_handler.DeleteHandler)
		cards.PATCH("/:id", cards_handler.PatchHandler)
		cards.POST("/", cards_handler.PostHandler)
	}
}

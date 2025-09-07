package collections_handlers

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostHandler(c *gin.Context) {
	req := Collection{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response_handlers.BadRequest(c, err)
		return
	}

	collection := schemas.Collection{
		CollectionId: uuid.New().String(),
		Name:         req.Name,
		LaunchYear:   req.LaunchYear,
		Type:         req.Type,
	}

	if err := db.Create(&collection).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"message":      "Collection created successfully",
		"collectionId": collection.CollectionId,
	}
	response_handlers.Created(c, msg)

}

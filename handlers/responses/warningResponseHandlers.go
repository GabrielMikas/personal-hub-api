package response_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context, err error) {
	mappedDetails := gin.H{
		"error": err.Error(),
	}
	mappedErr := gin.H{
		"detail": mappedDetails,
		"error":  "mrn:personalhub:error:notfound",
	}
	c.JSON(http.StatusNotFound, mappedErr)
}

func BadRequest(c *gin.Context, err error) {
	mappedDetails := gin.H{
		"error": err.Error(),
	}
	mappedErr := gin.H{
		"detail": mappedDetails,
		"error":  "mrn:personalhub:error:badrequest",
	}
	c.JSON(http.StatusBadRequest, mappedErr)
}

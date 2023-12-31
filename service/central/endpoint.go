package central

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct{}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) Health(c *gin.Context) {
	response := HealthResponse{
		Code:        200,
		Status:      "OK",
		Description: nil,
		Message:     nil,
	}
	c.JSON(http.StatusOK, response)
}

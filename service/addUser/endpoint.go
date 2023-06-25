package adduser

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct{}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) AddUser(c *gin.Context) {
	userService := NewUserService()
	var requestBody AddUserRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	if err := userService.AddUser(requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)
	} else {
		c.JSON(http.StatusCreated, nil)
	}
}

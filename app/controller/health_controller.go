package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Health(c *gin.Context)
}

type HealthControllerImpl struct {
}

func (u HealthControllerImpl) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func HealthControllerInit() *HealthControllerImpl {
	return &HealthControllerImpl{}
}

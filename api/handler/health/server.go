package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	HealthCheck(c *gin.Context)
}

type handler struct{}

func New() Handler {
	return &handler{}
}

func (h *handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

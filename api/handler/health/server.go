package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	Check(c *gin.Context)
}

type handler struct{}

func New() Handler {
	return &handler{}
}

func (h *handler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

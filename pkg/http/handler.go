package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Healthcheck(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

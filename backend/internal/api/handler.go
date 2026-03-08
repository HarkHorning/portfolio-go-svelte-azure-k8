package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// Add dependencies here (db, storage client, etc.)
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DevPrint(c *gin.Context) {
	c.String(http.StatusOK, "Hello, This endpoint is for development and does nothing. Good day!")
}

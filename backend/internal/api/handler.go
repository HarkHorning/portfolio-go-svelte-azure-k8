package api

import (
	"net/http"

	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/models"
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

func (h *Handler) DevArtModels(c *gin.Context) {
	tiles := []*models.ArtModel{
		models.CreateArtModel(1, "Painting 1", "/img/woman.jpg", "", true),
		models.CreateArtModel(2, "Painting 3", "/img/boat.jpg", "", false),
		models.CreateArtModel(3, "Painting 4", "/img/horse.jpg", "", true),
		models.CreateArtModel(4, "Painting 5", "/img/woman.jpg", "", true),
		models.CreateArtModel(5, "Painting 2", "/img/woman.jpg", "", true),
		models.CreateArtModel(6, "Painting 6", "/img/woman.jpg", "", true),
		models.CreateArtModel(7, "Painting 7", "/img/woman.jpg", "", true),
		models.CreateArtModel(8, "Painting 8", "/img/woman.jpg", "", true),
		models.CreateArtModel(9, "Painting 9", "/img/woman.jpg", "", true),
		models.CreateArtModel(10, "Painting 10", "/img/boat.jpg", "", false),
	}

	c.JSON(http.StatusOK, tiles)
}

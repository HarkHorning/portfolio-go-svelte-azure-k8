package api

import (
	"net/http"

	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/models"
	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/repo"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	sqlResource repo.Repo
}

func NewHandler(sqlResource repo.Repo) *Handler {
	return &Handler{
		sqlResource: sqlResource,
	}
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
		models.CreateArtModel(1, "Painting 1", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(2, "Boat on Lake", "https://artportfolio.blob.core.windows.net/lowgrade/boat.jpg", "", false),
		models.CreateArtModel(3, "Horse Watercolor", "https://artportfolio.blob.core.windows.net/lowgrade/horse.jpg", "", true),
		models.CreateArtModel(4, "Painting 5", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(6, "Painting 6", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(5, "Painting 2", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(7, "Horse Watercolor", "https://artportfolio.blob.core.windows.net/lowgrade/horse.jpg", "", true),
		models.CreateArtModel(8, "Painting 8", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(9, "Painting 9", "https://artportfolio.blob.core.windows.net/lowgrade/woman.jpg", "", true),
		models.CreateArtModel(10, "Boat on Lake", "https://artportfolio.blob.core.windows.net/lowgrade/boat.jpg", "", false),
	}

	c.JSON(http.StatusOK, tiles)
}

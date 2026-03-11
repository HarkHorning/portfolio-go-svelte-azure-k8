package api

import (
	"os"

	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routes(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	// Allows frontend to make requests to this API
	config := cors.Config{
		AllowOrigins:     getAllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // come back to this later
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	sqlResource := repo.NewRepo(db)

	handle := NewHandler(*sqlResource)

	router.GET("/health", handle.HealthCheck)

	art := router.Group("/api/art")
	{
		art.GET("/", handle.GetArtTiles)
	}

	return router
}

// getAllowedOrigins returns CORS origins based on environment
func getAllowedOrigins() []string {
	// Check for environment variable first (for production/docker)
	if origin := os.Getenv("CORS_ORIGIN"); origin != "" {
		return []string{origin}
	}

	// Default: allow local development origins
	return []string{
		"http://localhost:3000", // Svelte dev server
		"http://localhost:5173", // Vite dev server (alternative port)
	}
}

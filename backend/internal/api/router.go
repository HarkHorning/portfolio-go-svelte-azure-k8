package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	// CORS configuration
	// Allows frontend to make requests to this API
	config := cors.Config{
		AllowOrigins:     getAllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	handle := NewHandler()

	router.GET("/health", handle.HealthCheck)

	art := router.Group("/api/art")
	{
		art.GET("/", handle.DevArtModels)
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

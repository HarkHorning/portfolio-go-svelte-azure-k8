package api

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	handle := NewHandler()

	router.GET("/health", handle.HealthCheck)

	art := router.Group("/api/art")
	{
		art.GET("/", handle.DevPrint)
	}

	return router
}

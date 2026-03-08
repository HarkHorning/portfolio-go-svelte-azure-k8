package main

import (
	"log"

	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/api"
)

func main() {
	router := api.Routes()

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

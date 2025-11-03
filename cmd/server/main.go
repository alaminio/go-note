package main

import (
	"go-note/configs"
	"go-note/internal/api"
	"go-note/internal/db"
	"log"
)

func main() {
	// Load configuration
	configs.LoadConfig()
	config := configs.GetConfig()

	// Connect to database
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.CloseDB()

	// Setup router and start server
	router := api.SetupRouter(database)
	serverAddr := config.Host + ":" + config.Port
	log.Printf("Server starting on %s in %s mode", serverAddr, config.Environment)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

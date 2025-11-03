package main

import (
	"go-note/internal/api"
	"go-note/internal/db"
	"log"
)

func main() {
	// Connect to database
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.CloseDB()

	// Setup router and start server
	router := api.SetupRouter(database)
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

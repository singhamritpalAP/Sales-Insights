package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sales/internal/constants"
	"sales/internal/database"
	"sales/internal/handlers"
	"sales/pkg/cronjob"
)

func main() {

	db, err := database.NewDatabase(constants.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// Middleware to handle panics and recover
	router.Use(gin.Recovery())
	// Middleware for logging requests
	router.Use(gin.Logger())

	handlers.SetupRoutes(router, db)

	// Set up cron job in background
	go cronjob.SetupCronJob(db)

	if err := router.Run(constants.APIServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

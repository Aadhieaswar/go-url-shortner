package main

import (
	"log"
	"os"
	"path/filepath"
	"url-shortner/db"
	"url-shortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	logFile, err := os.OpenFile(filepath.Join("logs", "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("%v Failed to open log file %v", "S-999", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	router := gin.Default()

	database := db.SetupDatabase()
	db.Migrate(database)

	handlers.RegisterRoutes(router, database)

	router.Run(":8080")
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FranklynSistemas/chronofy/database"
	"github.com/FranklynSistemas/chronofy/internal/handlers"
	"github.com/FranklynSistemas/chronofy/internal/models"
	providers "github.com/FranklynSistemas/chronofy/internal/providers"
	"github.com/FranklynSistemas/chronofy/internal/repository"
	services "github.com/FranklynSistemas/chronofy/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Get the current environment (default to "local")
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// Load the corresponding .env file
	envFile := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}

	// Set Gin mode based on the GIN_MODE environment variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug" // Default to debug if not set
	}
	gin.SetMode(ginMode)

	fmt.Printf("Environment: %s\n", env)
	fmt.Printf("Gin mode is set to: %s\n", gin.Mode())
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.1"}) // Replace with your trusted proxy IPs

	DATABASE_URL := os.Getenv("DATABASE_URL")
	fmt.Printf("Database URL: %s\n", DATABASE_URL)
	repo, err := database.NewPostgresRepository(DATABASE_URL)
	if err != nil {
		log.Fatalf("Error creating postgres repository: %v", err)
	}

	repository.SetRepository(repo)

	// Define the GET endpoint.
	router.GET("/fetch-data", handlers.FetchDataHandler)

	// Start the server.
	router.Run(":8080") // Runs on localhost:8080 by default.
}

func testLocally() {
	// Define query parameters.
	params := models.QueryParams{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
		Filters:   map[string]string{"level": "error"},
		Order:     "asc",
	}

	// Initialize providers.
	providersList := []providers.Provider{
		&providers.GCPLogsProvider{},
		&providers.DatabaseProvider{},
		&providers.SentryProvider{},
	}

	// Create a context.
	ctx := context.Background()

	// Fetch data using the fetcher service.
	data, err := services.FetchDataFromProviders(ctx, params, providersList)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Print normalized results.
	fmt.Println("Normalized Data:")
	for _, d := range data {
		fmt.Printf("ID: %s, Source: %s, Timestamp: %s, Type: %s, Body: %v\n", d.ID, d.Source, d.Timestamp, d.Type, d.Body)
	}
}

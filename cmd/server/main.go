package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/somphonee/mlm-api/internal/app"
	"github.com/gofiber/fiber/v2/middleware/errors"
	"github.com/somphonee/mlm-api/config"
	"github.com/somphonee/mlm-api/pkg/database"
)
func main() {
	// Load config
	cfg := config.LoadConfig()
	
	// Connect to database
	db, err := database.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: app.ErrorHandler,
	})
	
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	
	// Setup routes
	app.SetupRoutes(app, db)
	
	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
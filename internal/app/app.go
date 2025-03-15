package app


import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/somphonee/mlm-api/config"
	"github.com/somphonee/mlm-api/internal/constants"
	"github.com/somphonee/mlm-api/pkg/database"
)

// App represents the application
type App struct {
	Fiber *fiber.App
	Config *config.Config
}

// New creates a new application instance
func New() *App {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      constants.AppName,
		ErrorHandler: errorHandler,
		BodyLimit:    constants.MaxBodySize, // 10MB
	})

	// Register middlewares
	registerMiddlewares(app)

	return &App{
		Fiber:  app,
		Config: cfg,
	}
}

// SetupRoutes configures all application routes
func (a *App) SetupRoutes() {
	// API version group
	api := a.Fiber.Group("/api/v1")

	// Health check endpoint
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"version": constants.AppVersion,
		})
	})

	// TODO: Register your routes here
	// Example:
	// auth.SetupRoutes(api)
	// users.SetupRoutes(api)
	// products.SetupRoutes(api)
	// orders.SetupRoutes(api)
	// commissions.SetupRoutes(api)
}

// Start starts the application server
func (a *App) Start() {
	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		_ = a.Fiber.Shutdown()
		database.CloseDB()
		log.Println("Application stopped")
		os.Exit(0)
	}()

	// Start the server
	addr := fmt.Sprintf(":%s", a.Config.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := a.Fiber.Listen(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// registerMiddlewares sets up all the application middlewares
func registerMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
			code = e.Code
	}
	return ctx.Status(code).JSON(fiber.Map{
			"error": err.Error(),
	})
}

// errorHandler is a custom error handler for Fiber
func errorHandler(c *fiber.Ctx, err error) error {
	// Default status code is 500
	code := fiber.StatusInternalServerError

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send error response
	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
		"code":    code,
	})
}

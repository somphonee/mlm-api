package database


import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	
	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	// Get database connection string from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/mlm_system"
		log.Println("Warning: Using default database URL. Set DATABASE_URL environment variable for production.")
	}

	// Configure connection pool
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("unable to parse database URL: %v", err)
	}

	// Set some reasonable pool settings
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	// Connect to the database with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	// Test the connection
	err = DB.Ping(ctx)
	if err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}

	log.Println("Connected to database successfully")
	return nil
}

// CloseDB gracefully closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}

// GetDB provides access to the database pool
func GetDB() *pgxpool.Pool {
	return DB
}
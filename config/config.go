package config
import (
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	JWTSecret   string
	JWTExpiry   time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	config := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/mlm_system"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiry:   time.Duration(getEnvAsInt("JWT_EXPIRY", 24)) * time.Hour,
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
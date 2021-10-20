package config

import (
	"github.com/joho/godotenv"
	"github.com/they-consulting/gogogo/console"
	"github.com/they-consulting/gogogo/helpers"
	"os"
)

func init() {
	envPaths := helpers.EnvPath()
	if err := godotenv.Load(envPaths[1], envPaths[0]); err != nil {
		console.Logger.Println("No env file found. Continuing.")
	}
}

type Config struct {
	Port                         string
}

type Database struct {
	Port         string
}

var conf Config

func New() {
	conf.Port = getEnvAsString("PORT", "3000")
}

func Get() *Config {
	return &conf
}

func getEnvAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

package config

import (
	"fmt"
	"os"
)

//This is a file for all my environment variable, basically much cleaner

// Config has all my api server values I could need
type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

// Holds the initial configuration once instead of initializing everytime
var Envs Config = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "Root1234"),
		DBAddress:  fmt.Sprintf("%s%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", ":3306")),
		DBName:     getEnv("DB_NAME", "roomieranks"),
	}
}

// Gets all our variables straight from our env variables
func getEnv(key, fallback string) string {
	//Attempt to first get our environment variable and if its is found return it
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	//Nothing might be found so we need some basic value in case it does not work
	return fallback
}

package config

import (
	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
	"log"
)

// Defines the name of the configuration file
const CONFIG_FILE = "config/application.yaml"

// Defines the application configuration
type AppConfig struct {
	// GIN mode
	GinMode string `json:"gin-gonic.gin.mode"`
	// Database name
	DatabaseName string `json:"db.name"`
	// Mapbox access token
	MapboxAccessToken string `json:"mapbox.access_token"`
}

// Read the application configuration properties
func ReadConfiguration() AppConfig {
	// Creates a config object
	var config AppConfig
	// Use a default name for the database
	config.DatabaseName = "app.db"
	// Use GIN release mode as default
	config.GinMode = gin.ReleaseMode
	// Reads the application configuration file
	if err := gonfig.GetConf(CONFIG_FILE, &config); err != nil {
		// If not, stop the application
		panic("Failure to read the configuration file!")
	}
	// Writes to the log whether Mapbox access is enabled or not
	log.Printf("Mapbox Enable: %t\n", len(config.MapboxAccessToken) != 0)
	// Writes the database filename to the log
	log.Printf("Database name: %s\n", config.DatabaseName)
	// Writes the GIN execution mode to the log
	log.Printf("GIN Mode: %s\n", config.GinMode)
	// Returns the application configuration values
	return config
}
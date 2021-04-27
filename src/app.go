package main

import (
	"roycetechnology.com/floverde/sample-rest-api/clients/mapbox"
    "roycetechnology.com/floverde/sample-rest-api/controllers"
	"roycetechnology.com/floverde/sample-rest-api/services"
	"roycetechnology.com/floverde/sample-rest-api/config"
    "roycetechnology.com/floverde/sample-rest-api/models"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Defines the application context.
//
// It keeps the reference to the state of the application.
type App struct {
    // GIN Engine
	router *gin.Engine
	// Database reference
	db     *gorm.DB
}

// Initializes the application context.
//
// It configures and prepares the components that make up the application.
// Specifically, it opens a connection to the database, creates the application
// components (REST controller and services) and registers the REST endpoints.
func (a *App) Initialize() {
	// Gets the defautl GIN engine
	a.router = gin.Default()
	// Fetch the application configuration properties
	config := config.ReadConfiguration()
    // Connect to the local database
	a.db = models.ConnectDatabase(config.DatabaseName)
	// Sets the GIN execution mode
	gin.SetMode(config.GinMode)
	
	// Creates a UserService using that database
	s := services.NewUserService(a.db)
	// #########################################
	mc := mapbox.NewMapboxClient(config.MapboxAccessToken)
	// Creates a UserRestController using that service
	c := controllers.NewUserRestController(s, mc)
	// Registers the REST endpoints of this controller
	c.RegisterHandlers(a.router)
}

// Run the application
//
// Starts each previously configured component.
func (a *App) Run() {
    // Run the server
    a.router.Run()
}
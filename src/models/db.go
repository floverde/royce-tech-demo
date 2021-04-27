package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Opens a connection to the local SQLite database.
// 
// The data stored in it remain available even when the server is restarted.
func ConnectDatabase(filename string) *gorm.DB {
    // Opens a connection to a SQLite3 database
	database, err := gorm.Open("sqlite3", filename)
	// Check that the operation was successful
	if err != nil {
		// If not, stop the application
		panic("Failed to connect to database!")
	}
	// Initialises the database structure to host user entities
	database.AutoMigrate(&User{})
	// Returns the database reference
	return database
}

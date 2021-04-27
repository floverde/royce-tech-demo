package models

import "time"

// Defines the data structure of a user stored on the database.
// 
// This data structure is also used as output by the REST read endpoints.
type User struct {
	Id           uint      `json:"id" gorm:"primary_key"`                      // user ID (must be unique)
	Name         string    `json:"name"`                                       // user name
	Dob          time.Time `json:"dob" time_format:"2006-01-02" time_utc:"1"`  // date of birth
	Address      string    `json:"address"`                                    // user address
	Description  string    `json:"description"`                                // user description
	CreatedAt    time.Time `json:"createdAt"`                                  // user created date
	UpdatedAt    time.Time `json:"updatedAt"`                                  // user updated date
}

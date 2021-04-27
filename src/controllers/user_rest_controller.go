package controllers

import (
	"roycetechnology.com/floverde/sample-rest-api/clients/mapbox"
    "roycetechnology.com/floverde/sample-rest-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// REST controller to manage users
type userRestController struct {
    service services.UserService
	mapboxClient mapbox.MapboxClient
}

// Creates and initialises the user REST controller
func NewUserRestController(service services.UserService, mapboxClient mapbox.MapboxClient) RestController {
    return &userRestController{service, mapboxClient}
}

// Registers the REST endpoints of the controller
func (rc *userRestController) RegisterHandlers(router *gin.Engine) {
    // Register endpoint to get all users
    router.GET("/users", rc.findUsers)
	// Register the endpoint for a single user
	router.GET("/users/:id", rc.findUser)
	// Register the endpoint to create a new user
	router.POST("/users", rc.createUser)
	// Register the endpoint to update a user
	router.PATCH("/users/:id", rc.updateUser)
	// Register the endpoint to delete a user
	router.DELETE("/users/:id", rc.deleteUser)
	// Register endpoint to get the places
	// associated with the user's address
	router.GET("/users/:id/places", rc.getUserPlaces)
}

// GET /users
//
// Find all users
func (rc *userRestController) findUsers(cxt *gin.Context) {
    // Invokes the service by getting all users
	users, err := rc.service.GetAll()
	// Check that the previous operation was successful
	if err != nil {
		// Returns an error 400 BAD REQUEST attaching the error message
        cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}
	
	// Formats the list of users in JSON
	cxt.JSON(http.StatusOK, users)
}

// GET /users/:id
//
// Find a single user
func (rc *userRestController) findUser(cxt *gin.Context) {
	// Declares a unsigned integer value
    var id uint
	// Retrieves the requested user ID from the request URL
	if fetchUintURLParam(cxt, "id", &id) {
	
		// Invokes the service to get the user by ID
		user, err := rc.service.GetById(id)
		// Check that the previous operation was successful
		if err != nil {
			// Check if a NOT FOUND error has been raised
			if (strings.Contains(err.Error(), "not found")) {
				// Returns an error 404 NOT FOUND attaching an error description
				cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			} else {
				// Returns an error 400 BAD REQUEST attaching the error message
				cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			return
		}
		
		// Checks whether the user exists within the repository
		if user == nil {
			// Returns an error 404 NOT FOUND attaching an error description
			cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			return
		}
		
		// Formats user data in JSON
		cxt.JSON(http.StatusOK, user)
	}
}

// POST /users
//
// Create new user
func (rc *userRestController) createUser(cxt *gin.Context) {
	// Declares the DTO for the creation of the user
	var input services.UserInputDTO
	// Parses the body of the JSON request by valuing the DTO
	if err := cxt.ShouldBindJSON(&input); err != nil {
		// Returns an error 400 BAD REQUEST attaching the error message
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // Invokes the service to create a new user
    user, err := rc.service.Create(input)
	// Check that the previous operation was successful
	if err != nil {
		// Returns an error 400 BAD REQUEST attaching the error message
        cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}

	// Formats user data in JSON
	cxt.JSON(http.StatusCreated, user)
}

// PATCH /users/:id
//
// Update a user
func (rc *userRestController) updateUser(cxt *gin.Context) {
    // Declares a unsigned integer value
    var id uint
	// Retrieves the requested user ID from the request URL
	if fetchUintURLParam(cxt, "id", &id) {

		// Declares the DTO for the creation of the user
		var input services.UserInputDTO
		// Parses the body of the JSON request by valuing the DTO
		if err := cxt.ShouldBindJSON(&input); err != nil {
			// Returns an error 400 BAD REQUEST attaching the error message
			cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Invokes the service to update the user's data
		if user, err := rc.service.Update(id, input); err != nil {
			// Check if a NOT FOUND error has been raised
			if strings.Contains(err.Error(), "not found") {
				// Returns an error 404 NOT FOUND attaching an error description
				cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			} else {
				// Returns an error 400 BAD REQUEST attaching the error message
				cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
		} else {
			// Formats user data in JSON
			cxt.JSON(http.StatusOK, user)
		}
	}
}

// DELETE /users/:id
//
// Delete a user
func (rc *userRestController) deleteUser(cxt *gin.Context) {
	// Declares a unsigned integer value
    var id uint
	// Retrieves the requested user ID from the request URL
	if fetchUintURLParam(cxt, "id", &id) {
		// Invokes the service to delete a user
		if err := rc.service.Delete(id); err != nil {
			// Check if a NOT FOUND error has been raised
			if strings.Contains(err.Error(), "not found") {
				// Returns an error 404 NOT FOUND attaching an error description
				cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			} else {
				// Returns an error 400 BAD REQUEST attaching the error message
				cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
		} else {
			// Returns an empty message
			// with status HTTP 200 OK
			cxt.Status(http.StatusOK)
		}
	}
}

// GET /users/:id/places
//
// Gets the list of places associated with the user's address
func (rc *userRestController) getUserPlaces(cxt *gin.Context) {
	// Declares a unsigned integer value
    var id uint
	// Retrieves the requested user ID from the request URL
	if fetchUintURLParam(cxt, "id", &id) {
	
		// Invokes the service to get the user by ID
		user, err := rc.service.GetById(id)
		
		// Check that the previous operation was successful
		if err != nil {
			// Check if a NOT FOUND error has been raised
			if strings.Contains(err.Error(), "not found") {
				// Returns an error 404 NOT FOUND attaching an error description
				cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			} else {
				// Returns an error 400 BAD REQUEST attaching the error message
				cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			return
		}
		
		// Checks whether the user exists within the repository
		if user == nil {
			// Returns an error 404 NOT FOUND attaching an error description
			cxt.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			return
		}
		
		// Checks whether the user has an address
		if len(user.Address) > 0 {
			// Check if the Mapbox client is available
			if rc.mapboxClient != nil {
				// Gets the list of places associated with the user's address
				response, err := rc.mapboxClient.GetPlaces(user.Address)
				// Check that the previous operation was successful
				if err != nil {
					// Returns an error 400 BAD REQUEST attaching the error message
					cxt.JSON(http.StatusBadRequest, gin.H{"error": err})
					return
				}
				// Returns the contents of the Mapbox response
				cxt.DataFromReader(http.StatusOK, response.ContentLength,
								   gin.MIMEJSON, response.Body, nil)
			} else {
				// Returns only the user's address in JSON format
				cxt.JSON(http.StatusOK, gin.H{"address": user.Address})
			}
		} else {
			// Returns an empty message with
			// status HTTP 204 NO CONTENT
			cxt.Status(http.StatusNoContent)
		}
	}
}
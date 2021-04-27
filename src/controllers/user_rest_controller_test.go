package controllers_test

import (
    "roycetechnology.com/floverde/sample-rest-api/controllers"
	"roycetechnology.com/floverde/sample-rest-api/services"
	
	"github.com/gin-gonic/gin"
	"encoding/json"
	
	"net/http/httptest"
	"net/http"
	
	"strings"
	"testing"
	"bytes"

	"fmt"
)

// TestUserController - entry point
func TestUserController(t *testing.T) {
    // Gets the defautl GIN engine
	r := gin.Default()
	// Create a stub instance of UserService
    s := services.NewUserServiceStub()
	// Creates a UserRestController using the stub service
	c := controllers.NewUserRestController(s, nil)
	// Registers the REST endpoints of this controller
	c.RegisterHandlers(r)
	
	// 01 - Test: emptyRepository
	testEmptyRepository(r, t)
	// 02 - Test: createUser
	testCreateUser(r, t)
	// 03 - Test: findSingleUser
	testFindSingleUser(r, t)
	// 04 - Test: findNotExistingUser
	testFindNotExistingUser(r, t)
	// 05 - Test: updateUser
	testUpdateUser(r, t)
	// 06 - Test: deleteUser
	testDeleteUser(r, t)
}

/*************************************************************/
/**************** INTERNAL UTILITY METHODS *******************/
/*************************************************************/

// [Internal Utility Method]: Executes an HTTP request, retrieving its response.
func executeRequest(r *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

// [Internal Utility Method]: Checks that the HTTP status of the response is as expected.
func checkResponseCode(t *testing.T, expected, actual int) bool {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
		return false
	}
	return true
}

/*************************************************************/
/******************** TESTING METHODS ************************/
/*************************************************************/
// 01 - Test: emptyRepository
//
// Check that the repository is initially empty. It uses the method
// to retrieve the list of all users and checks that this list is empty.
func testEmptyRepository(r *gin.Engine, t *testing.T) {
    t.Run("emptyRepository", func(t *testing.T) {
	    // [Step 1 - PREPARE HTTP REQUEST]
		// Creates an HTTP request to get all users
		req, err := http.NewRequest("GET", "/users", nil)
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 200 OK
		if checkResponseCode(t, http.StatusOK, response.Code) {
		
			// [Step 3 - CHECK]
			// Check that the body of the HTTP response is empty
			if body := response.Body.String(); body != "[]" {
				t.Errorf("Expected an empty array. Got %s", body)
			}
		}
	})
}

// 02 - Test: createUser
//
// Create a user by checking that this user is correctly stored in the repository.
func testCreateUser(r *gin.Engine, t *testing.T) {
    t.Run("createUser", func(t *testing.T) {
	    // Declares the JSON string of user creation parameters
	    var jsonStr = bytes.NewBuffer([]byte(`{"name":"Giggi"}`))
		
		// [Step 1 - PREPARE HTTP REQUEST]
		// Creates the HTTP request to create a new user
		req, err := http.NewRequest("POST", "/users", jsonStr)
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		// Sets the HTTP header that specifies the format of the request
		req.Header.Set("Content-Type", "application/json")
	
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 201 CREATED
		if checkResponseCode(t, http.StatusCreated, response.Code) {
		
			// [Step 3 - FETCH]
			// Declares a map with string keys
			var m map[string]interface{}
			// Unmarshals the HTTP response body
			json.Unmarshal(response.Body.Bytes(), &m)

			// [Step 4 - CHECK]
			// Gets the numerical value of the user ID
			userId := uint(m["id"].(float64))
			// Check that the user ID has been assigned
			if userId == 0 {
				t.Error("User Id not assigned")
				return
			}

			// [Step 5 - CHECK]
			// Check if the user name is as expected
			if m["name"] != "Giggi" {
				t.Errorf("Expected user name to be 'Giggi'. Got '%v'", m["name"])
				return
			}
			
			// [Step 6 - PREPARE HTTP REQUEST]
			// Builds the URL to retrieve the newly created user
			reqUrl := fmt.Sprintf("/users/%d", userId)
			// Prepare the HTTP request to retrieve the newly created user
			req, err = http.NewRequest("GET", reqUrl, nil)
			// Check that the previous operation was successful
			if err != nil {
				t.Error(err)
				return
			}
			
			// [Step 7 - EXECUTE + CHECK]
			// Execute the previous built HTTP request
			response := executeRequest(r, req)
			// Check that the HTTP response status is 200 OK
			checkResponseCode(t, http.StatusOK, response.Code)
		}
	})
}

// 03 - Test: findSingleUser
//
// Check that the search by ID works properly
func testFindSingleUser(r *gin.Engine, t *testing.T) {
    t.Run("findSingleUser", func(t *testing.T) {
		// [Step 1 - PREPARE HTTP REQUEST]
		// Prepares the HTTP request to get a single user
		req, err := http.NewRequest("GET", "/users/1", nil)
		// Check that the previous operation was successful
		if err != nil {
			t.Error(err)
			return
		}
		
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 200 OK
		if checkResponseCode(t, http.StatusOK, response.Code) {
		
			// [Step 3 - FETCH]
			// Declares a map with string keys
			var m map[string]interface{}
			// Unmarshals the HTTP response body
			json.Unmarshal(response.Body.Bytes(), &m)
			
			// [Step 4 - CHECK]
			// Gets the numerical value of the user ID
			// and check that it is the expected value
			if userId := uint(m["id"].(float64)); userId != 1 {
				t.Errorf("Expected ID must be 1. Got %d", userId)
			}
		}
	})
}

// 04 - Test: findNotExistingUser
//
// Check that searching by ID for a non-existent user does not produce any results
func testFindNotExistingUser(r *gin.Engine, t *testing.T) {
    t.Run("findNotExistingUser", func(t *testing.T) {
		// [Step 1 - PREPARE HTTP REQUEST]
		// Prepares the HTTP request to get a user that does not exist
		req, err := http.NewRequest("GET", "/users/999", nil)
		// Check that the previous operation was successful
		if err != nil {
			t.Error(err)
			return
		}
		
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 404 NOT FOUND
		if checkResponseCode(t, http.StatusNotFound, response.Code) {

			// [Step 3 - FETCH]
			// Declares a map with string keys
			var m map[string]interface{}
			// Unmarshals the HTTP response body
			json.Unmarshal(response.Body.Bytes(), &m)
			
			// [Step 4 - CHECK]
			// Check that the error message contains the "not found" substring
			if !strings.Contains(m["error"].(string), "not found") {
				t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
			}
		}
	})
}

// 05 - Test: updateUser
//
// Checks that the updating of user data is working
func testUpdateUser(r *gin.Engine, t *testing.T) {
    t.Run("updateUser", func(t *testing.T) {
		// Declares the JSON string of a user's update parameters
		var jsonStr = bytes.NewBuffer([]byte(`{"name":"Giovanni"}`))
		
		// [Step 1 - PREPARE HTTP REQUEST]
		// Prepares the HTTP request to update a user's data
		req, err := http.NewRequest("PATCH", "/users/1", jsonStr)
		// Check that the previous operation was successful
		if err != nil {
			t.Error(err)
			return
		}
		
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 200 OK
		if checkResponseCode(t, http.StatusOK, response.Code) {
		
			// [Step 3 - FETCH]
			// Declares a map with string keys
			var m map[string]interface{}
			// Unmarshals the HTTP response body
			json.Unmarshal(response.Body.Bytes(), &m)
			
			// [Step 4 - CHECK]
			// Gets the numerical value of the user ID
			// and check that it is the expected value
			if userId := uint(m["id"].(float64)); userId != 1 {
				t.Errorf("Expected the id to remain the same (1). Got %d", userId)
			}
			
			// [Step 5 - CHECK]
			// Check if the user name is as expected
			if m["name"] != "Giovanni" {
				t.Errorf("'Name' field not updated (expected: Giovanni, found: %s)", m["name"])
			}
		}
	})
}

// 06 - Test: deleteUser
//
// Checks that the procedure for deleting a user works
func testDeleteUser(r *gin.Engine, t *testing.T) {
    t.Run("deleteUser", func(t *testing.T) {
		// [Step 1 - PREPARE HTTP REQUEST]
		// Prepares the HTTP request to delete a user
		req, err := http.NewRequest("DELETE", "/users/1", nil)
		// Check that the previous operation was successful
		if err != nil {
			t.Error(err)
			return
		}
		
		// [Step 2 - EXECUTE + CHECK]
		// Execute the previous built HTTP request
		response := executeRequest(r, req)
		// Check that the HTTP response status is 200 OK
		if checkResponseCode(t, http.StatusOK, response.Code) {
		
			// [Step 3 - PREPARE HTTP REQUEST]
			// Prepares the HTTP request to get the previously deleted user
			req, err = http.NewRequest("GET", "/users/1", nil)
			// Check that the previous operation was successful
			if err != nil {
				t.Error(err)
				return
			}
			
			// [Step 2 - EXECUTE + CHECK]
			// Execute the previous built HTTP request
			response := executeRequest(r, req)
			// Check that the HTTP response status is 404 NOT FOUND
			// proving that the user is no longer contained in the repository
			checkResponseCode(t, http.StatusNotFound, response.Code)
		}
	})
}
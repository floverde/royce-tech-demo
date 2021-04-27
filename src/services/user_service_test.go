package services_test

import (
	"roycetechnology.com/floverde/sample-rest-api/services"
	"roycetechnology.com/floverde/sample-rest-api/models"
	"strings"
	"testing"
)

// Defines the SQLite URL to open an in-memory database
const IN_MEMORY_DATABASE = "file::memory:?cache=shared"

// TestUserService - entry point
func TestUserService(t *testing.T) {
	// Opens a connection with an in-memory database
	db := models.ConnectDatabase(IN_MEMORY_DATABASE)
	// Create a new UserSerivce instance
	s := services.NewUserService(db)
	
	// 01 - Test: emptyRepository
	testEmptyRepository(s, t)
	// 02 - Test: createUser
	testCreateUser(s, t)
	// 03 - Test: findSingleUser
	testFindSingleUser(s, t)
	// 04 - Test: findNotExistingUser
	testFindNotExistingUser(s, t)
	// 05 - Test: updateUser
	testUpdateUser(s, t)
	// 06 - Test: deleteUser
	testDeleteUser(s, t)
}

/*************************************************************/
/******************** TESTING METHODS ************************/
/*************************************************************/
// 01 - Test: emptyRepository
//
// Check that the repository is initially empty. It uses the method
// to retrieve the list of all users and checks that this list is empty.
func testEmptyRepository(s services.UserService, t *testing.T) {
	t.Run("emptyRepository", func(t *testing.T) {
		// [Step 1 - EXECUTE + ERROR CHECK]
		// Retrieve the list of all users
		users, err := s.GetAll()
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		
		// [Step 2 - CHECK]
		// Check that the list of users returned is empty
		if count := len(users); count != 0 {
			t.Errorf("Expected an empty array. Actually it contains %d items", count)
		}
	})
}

// 02 - Test: createUser
//
// Create a user by checking that this user is correctly stored in the repository.
func testCreateUser(s services.UserService, t *testing.T) {
	t.Run("createUser", func(t *testing.T) {
		// Declares the input parameters for creating a user
	    params := services.UserInputDTO{Name: "Giggi"}
	
		// [Step 1 - EXECUTE]
		// Execute the method to create a new user
		user, err := s.Create(params)
		
		// [Step 2 - CHECK]
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		
		// [Step 3 - CHECK]
		// Check that the user ID has been assigned
		if user.Id == 0 {
			t.Error("User Id not assigned")
			return
		}
		
		// [Step 4 - CHECK]
		// Check if the user name is as expected
		if user.Name != params.Name {
			t.Errorf("Expected %s as Name value. Found: %s", params.Name, user.Name)
			return
		}
		
		// [Step 5 - EXECUTE]
		// Requests the repository to provide the user it has just
		// created, to check that he or she is actually present in it
		user, err = s.GetById(user.Id)
		
		// [Step 6 - CHECK]
		// Check that the previous operation was successful
		// (note: the NOT FOUND error is an expected error)
		if err != nil && !strings.Contains(err.Error(), "not found") {
		    t.Error(err)
			return
		}
		
		// [Step 6 - CHECK]
		// Check that the user has been returned successfully
		if user == nil {
		    t.Error("User not stored")
			return
		}
	})
}

// 03 - Test: findSingleUser
//
// Check that the search by ID works properly
func testFindSingleUser(s services.UserService, t *testing.T) {
	t.Run("findSingleUser", func(t *testing.T) {
		// [Step 1 - EXECUTE + ERROR CHECK]
		// Retrieves the user created in the previous test (02 - createUser)
		user, err := s.GetById(1)
		// Check that the previous operation was successful
		// (note: the NOT FOUND error is an expected error)
		if err != nil && !strings.Contains(err.Error(), "not found") {
		    t.Error(err)
			return
		}
		
		// [Step 2 - CHECK]
		// Check that the user has been returned successfully
		if user == nil {
		    t.Error("Expected user with ID 1, but it was not found")
			return
		}
		
		// [Step 3 - CHECK]
		// Check that the user ID is the expected one
		if user.Id != 1 {
		    t.Errorf("Returned user has a different ID than expected (expected: 1, found %d)", user.Id)
			return
		}
	})
}

// 04 - Test: findNotExistingUser
//
// Check that searching by ID for a non-existent user does not produce any results
func testFindNotExistingUser(s services.UserService, t *testing.T) {
	t.Run("findNotExistingUser", func(t *testing.T) {
		// [Step 1 - EXECUTE]
		// Tries to get a user that does not exist
		user, err := s.GetById(999)
		
		// [Step 2 - CHECK]
		// Check that the previous operation was successful
		// (note: the NOT FOUND error is an expected error)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			t.Error(err)
			return
		}
		
		// [Step 3 - CHECK]
		// Checks that the user has not been returned
		if user != nil {
			t.Errorf("Unexpected user %v\n", user)
			return
		}
	})
}

// 05 - Test: updateUser
//
// Checks that the updating of user data is working
func testUpdateUser(s services.UserService, t *testing.T) {
	t.Run("updateUser", func(t *testing.T) {
		// Declares the input parameters for updating a user
		params := services.UserInputDTO{Name: "Giovanni"}
		
		// [Step 1 - EXECUTE]
		// Execute the method to update the user data
		user, err := s.Update(1, params)
		
		// [Step 2 - CHECK]
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		
		// [Step 3 - CHECK]
		// Check that the user has been returned successfully
		if user == nil {
		    t.Error("Expected user with ID 1, but it was not found")
			return
		}
		
		// [Step 4 - CHECK]
		// Check that the user ID is the expected one
		if user.Id != 1 {
		    t.Errorf("Returned user has a different ID than expected (expected: 1, found %d)", user.Id)
			return
		}
		
		// [Step 5 - CHECK]
		// Check if the user name is as expected
		if user.Name != params.Name {
			t.Errorf("Expected %s as Name value. Found: %s", params.Name, user.Name)
			return
		}
	})
}

// 06 - Test: deleteUser
//
// Checks that the procedure for deleting a user works
func testDeleteUser(s services.UserService, t *testing.T) {
	t.Run("deleteUser", func(t *testing.T) {
		
		// [Step 1 - EXECUTE]
		// Executes the method to delete a user
		err := s.Delete(1)
		
		// [Step 2 - CHECK]
		// Check that the previous operation was successful
		if err != nil {
		    t.Error(err)
			return
		}
		
		// [Step 3 - EXECUTE]
		// Tries to retrieve the deleted user
		user, err2 := s.GetById(1)
		
		// [Step 4 - CHECK]
		// Check that the previous operation was successful
		// (note: the NOT FOUND error is an expected error)
		if err2 != nil && !strings.Contains(err2.Error(), "not found") {
		    t.Error(err2)
			return
		}
		
		// [Step 5 - CHECK]
		// Checks that the user is no longer in the repository
		if user != nil {
		    t.Error("The user with ID 1 is still contained in the repository")
		}
	})
}
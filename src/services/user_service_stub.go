package services

import (
	"roycetechnology.com/floverde/sample-rest-api/models"
	"time"
)

type userServiceStub struct {
    users []models.User
	lastId uint
}

func NewUserServiceStub() UserService {
    return &userServiceStub{make([]models.User, 0, 10), 0}
}
    
// Find all users
func (s *userServiceStub) GetAll() ([]models.User, error) {
	return s.users, nil
}

// Find a single user
func (s *userServiceStub) GetById(id uint) (*models.User, error) {
	for _, user := range s.users {
	    if user.Id == id {
		    return &user, nil
		}
	}
	return nil, nil
}

// Create new user
func (s *userServiceStub) Create(params UserInputDTO) (*models.User, error) {
    s.lastId++
	
	// Create the user entity from the provided DTO
	user := models.User{Id: s.lastId,
                        Name: params.Name,
                        Dob: params.Dob,
						Address: params.Address,
						Description: params.Description,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now()}
	
	s.users = append(s.users, user)
	
	return &user, nil
}

// Update a user
func (s *userServiceStub) Update(id uint, params UserInputDTO) (*models.User, error) {
	user, err := s.GetById(id)
	
	if user != nil && err == nil {
		user.Name = params.Name
		user.Dob = params.Dob
		user.Address = params.Address
		user.Description = params.Description
		user.UpdatedAt = time.Now()
	}

	return user, err
}

// Delete a user
func (s *userServiceStub) Delete(id uint) error {
	// Get model if exist
	for i, user := range s.users {
	    if user.Id == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
		    return nil
		}
	}
	return nil
}

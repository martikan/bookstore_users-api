package service

import (
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/util/crypto_utils"
	"github.com/martikan/bookstore_users-api/util/date_utils"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

// Interface for User service
type userServiceInterface interface {
	Get(int64) (*user.User, *errors.RestError)
	Search(string) (user.Users, *errors.RestError)
	Create(user.User) (*user.User, *errors.RestError)
	Update(bool, user.User) (*user.User, *errors.RestError)
	Delete(int64) *errors.RestError
}

// Function for get a user by id
func (s *userService) Get(id int64) (*user.User, *errors.RestError) {

	result := &user.User{Id: id}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// Function for search a user
func (s *userService) Search(str string) (user.Users, *errors.RestError) {
	dao := &user.User{}
	return dao.FindByStatus(str)
}

// Function for create a user
func (s *userService) Create(u user.User) (*user.User, *errors.RestError) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.Status = user.StatusActive
	u.CreatedAt = date_utils.GetNowDBFormat()

	// Encrypt password
	passwordHash, hashingErr := crypto_utils.HashPassword(u.Password)
	if hashingErr != nil {
		return nil, hashingErr
	}
	u.Password = passwordHash

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

// Function for update a user by id
func (s *userService) Update(partial bool, u user.User) (*user.User, *errors.RestError) {

	currentUser, err := UserService.Get(u.Id)
	if err != nil {
		return nil, err
	}

	if partial {
		if u.FirstName != "" {
			currentUser.FirstName = u.FirstName
		}
		if u.LastName != "" {
			currentUser.LastName = u.LastName
		}
		if u.Email != "" {
			currentUser.Email = u.Email
		}
	} else {
		currentUser.FirstName = u.FirstName
		currentUser.LastName = u.LastName
		currentUser.Email = u.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

// Function for delete a user by id
func (s *userService) Delete(id int64) *errors.RestError {
	user := &user.User{Id: id}
	return user.Delete()
}

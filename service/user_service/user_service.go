package user_service

import (
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/error"
)

func GetUser(id int64) (*user.User, *error.RestError) {

	result := &user.User{Id: id}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// Function for create a user
func CreateUser(u user.User) (*user.User, *error.RestError) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

package user_service

import (
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/util/date_utils"
)

// Function for get a user by id
func Get(id int64) (*user.User, *errors.RestError) {

	result := &user.User{Id: id}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// Function for search a user
func Search(s string) ([]user.User, *errors.RestError) {
	dao := &user.User{}
	return dao.FindByStatus(s)
}

// Function for create a user
func Create(u user.User) (*user.User, *errors.RestError) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.Status = user.StatusActive
	u.CreatedAt = date_utils.GetNowDBFormat()

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

// Function for update a user by id
func Update(partial bool, u user.User) (*user.User, *errors.RestError) {

	currentUser, err := Get(u.Id)
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
func Delete(id int64) *errors.RestError {
	user := &user.User{Id: id}
	return user.Delete()
}

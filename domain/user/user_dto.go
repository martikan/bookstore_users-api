package user

import (
	"strings"
	"time"

	"github.com/martikan/bookstore_users-api/error"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Method for validate a user
func (u *User) Validate() *error.RestError {

	u.FirstName = strings.TrimSpace(strings.ToLower(u.FirstName))

	if u.FirstName == "" {
		return error.NewBadRequestError("Invalid first name.")
	}

	u.LastName = strings.TrimSpace(strings.ToLower(u.LastName))

	if u.LastName == "" {
		return error.NewBadRequestError("Invalid last name.")
	}

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return error.NewBadRequestError("Invalid email address.")
	}

	return nil
}

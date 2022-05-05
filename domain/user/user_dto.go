package user

import (
	"strings"

	"github.com/martikan/bookstore_users-api/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

// Validate Method for validate a user
func (u *User) Validate() *errors.RestError {

	u.FirstName = strings.TrimSpace(u.FirstName)
	if u.FirstName == "" {
		return errors.NewBadRequestError("First name is required.")
	}
	if len(u.FirstName) > 45 {
		return errors.NewBadRequestError("First name cannot be greater then 45 characters.")
	}

	u.LastName = strings.TrimSpace(u.LastName)
	if u.LastName == "" {
		return errors.NewBadRequestError("Last name is required.")
	}
	if len(u.LastName) > 45 {
		return errors.NewBadRequestError("Last name cannot be greater then 45 characters.")
	}

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Email address is required.")
	}
	if len(u.Email) > 255 {
		return errors.NewBadRequestError("Email address cannot be greater then 45 characters.")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("Password is required.")
	}
	if len(u.Password) > 66 {
		return errors.NewBadRequestError("Password cannot be greater then 66 characters.")
	}

	u.Status = strings.TrimSpace(strings.ToLower(u.Status))
	if u.Status != "" && u.Status != "active" && u.Status != "inactive" {
		return errors.NewBadRequestError("Status can be 'active' or 'inactive' only.")
	}

	return nil
}

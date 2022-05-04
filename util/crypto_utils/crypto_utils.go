package crypto_utils

import (
	"github.com/martikan/bookstore_users-api/errors"
	"golang.org/x/crypto/bcrypt"
)

// Hasing password with Bcrypt
func HashPassword(input string) (string, *errors.RestError) {

	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.NewInternalServerError(err.Error())
	}

	return string(hash), nil
}

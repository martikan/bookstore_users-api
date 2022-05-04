package crypto_utils

import (
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/logger"
	"golang.org/x/crypto/bcrypt"
)

// Hasing password with Bcrypt
func HashPassword(input string) (string, *errors.RestError) {

	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error when trying to hash the password", err)
		return "", errors.NewInternalServerError("Api error")
	}

	return string(hash), nil
}

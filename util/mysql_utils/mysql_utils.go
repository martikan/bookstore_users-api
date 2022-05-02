package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/martikan/bookstore_users-api/errors"
)

const (
	noRowsFoundError = "no rows in result set"
)

func ParseError(err error) *errors.RestError {

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsFoundError) {
			return errors.NewNotFoundError("No record found")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing database response: %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Email address is already exists.")
	}
	return errors.NewInternalServerError(fmt.Sprintf("Error processing response: %s", err.Error()))
}

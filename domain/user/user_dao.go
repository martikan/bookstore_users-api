package user

import (
	"fmt"
	"strings"

	"github.com/martikan/bookstore_users-api/datasource/mysql/users_db"
	"github.com/martikan/bookstore_users-api/error"
	"github.com/martikan/bookstore_users-api/util/date_utils"
)

const (
	tableName = "users"

	// Indexes

	uniqueEmailIndex = "email_UNIQUE"

	// Queries

	insertUser = "INSERT INTO " + tableName + " (first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"

	deleteUserById = "DELETE FROM" + tableName + " WHERE id = ?;"

	findUserById = "" +
		"SELECT " +
		"u.id, " +
		"u.first_name, " +
		"u.last_name, " +
		"u.email " +
		"FROM " + tableName + " u " +
		"WHERE id = ?;"
)

func (u *User) Get() *error.RestError {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	// TODO: Fix it!

	// result := usersDB[u.Id]

	// if result == nil {
	// 	return error.NewNotFoundError(fmt.Sprintf("User with id=%d not found.", u.Id))
	// }

	// u.Id = result.Id
	// u.FirstName = result.FirstName
	// u.LastName = result.LastName
	// u.Email = result.Email
	// u.CreatedAt = result.CreatedAt

	return nil
}

func (u *User) Save() *error.RestError {

	stmt, err := users_db.Client.Prepare(insertUser)
	if err != nil {
		return error.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	u.CreatedAt = date_utils.GetNow()

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), uniqueEmailIndex) {
			return error.NewBadRequestError("Email address is already exist")
		}
		return error.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	id, err := result.LastInsertId()
	if err != nil {
		return error.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	u.Id = id

	return nil
}

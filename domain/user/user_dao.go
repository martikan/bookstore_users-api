package user

import (
	"github.com/martikan/bookstore_users-api/datasource/mysql/users_db"
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/util/mysql_utils"
)

const (
	tableName = "USERS"

	// Queries

	insertUser = "INSERT INTO " + tableName +
		" (first_name, last_name, email, created_at, status, password) VALUES (?, ?, ?, ?, ?, ?);"

	updateUser = "UPDATE " + tableName +
		" SET" +
		" first_name = ?, last_name = ?, email = ?, status = ?, password = ?" +
		" WHERE id = ?;"

	deleteUserById = "DELETE FROM" + tableName +
		" WHERE id = ?;"

	// findAllUser = "" +
	// 	"SELECT " +
	// 	"id, " +
	// 	"first_name, " +
	// 	"last_name, " +
	// 	"email " +
	// 	"status " +
	//  "password " +
	// 	"FROM " + tableName + ";"

	findUserByStatus = "" +
		"SELECT " +
		"id, " +
		"first_name, " +
		"last_name, " +
		"email " +
		"status " +
		"password " +
		"FROM " + tableName +
		" WHERE status = ?;"

	findUserById = "" +
		"SELECT " +
		"id, " +
		"first_name, " +
		"last_name, " +
		"email " +
		"status " +
		"password " +
		"FROM " + tableName +
		" WHERE id = ?;"
)

func (u *User) Get() *errors.RestError {

	stmt, err := users_db.Client.Prepare(findUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.Password); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (u *User) FindByStatus(s string) ([]User, *errors.RestError) {

	stmt, err := users_db.Client.Prepare(findUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query(s)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &u.Password); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No users found.")
	}

	return results, nil
}

func (u *User) Update() *errors.RestError {

	stmt, err := users_db.Client.Prepare(updateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (u *User) Save() *errors.RestError {

	stmt, err := users_db.Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.CreatedAt, u.Status, u.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	u.Id = id

	return nil
}

func (u *User) Delete() *errors.RestError {

	stmt, err := users_db.Client.Prepare(deleteUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

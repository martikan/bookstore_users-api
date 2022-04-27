package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_db_username = "mysql_users_db_username"
	mysql_users_db_password = "mysql_users_db_password"
	mysql_users_db_uri      = "mysql_users_db_uri"
	mysql_users_db_database = "mysql_users_db_database"
)

var (
	Client *sql.DB

	db_username = os.Getenv(mysql_users_db_username)
	db_password = os.Getenv(mysql_users_db_password)
	db_uri      = os.Getenv(mysql_users_db_uri)
	db_database = os.Getenv(mysql_users_db_database)
)

// Create a connection to the database
func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		db_username, db_password, db_uri, db_database)

	var err error

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured.")

}

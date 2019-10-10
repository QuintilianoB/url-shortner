package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"metro-ag/util"
)

// Errors for database
var (
	DatabaseQueryError = util.ErrorMessage{"Database error. Please, contact the system administrators."}
)

func ConnectDB() *sql.DB {

	driver := util.GetEnv("dbDriver", "mysql")
	username := util.GetEnv("dbUser", "")
	password := util.GetEnv("dbPwd", "")
	host := util.GetEnv("dbHost", "127.0.0.1")
	port := util.GetEnv("dbPort", "3306")
	database := util.GetEnv("dbName", "")

	if username == "" || password == "" || database == "" {
		log.Fatal("Invalid database configuration. Check .env file.")
	}

	dbUrl := fmt.Sprintf(username + ":" + password + "@(" + host + ":" + port + ")/" + database)

	db, err := sql.Open(driver, dbUrl)
	util.LogFatal(err)

	return db
}

package store

import (
	"io/ioutil"
	"log"
	"metro-ag/util"
	"strings"
)

/*
	MysqlCreate will read the initial sql file and import it.
	The SQL files should be idempotent, eg. "create table if not exits."
	Its not safe for production environment and should be use only for development and/or testing.
 */
func MysqlCreate() {

	file, err := ioutil.ReadFile("store/sql/initial.sql")

	if err != nil {
		log.Fatal("Error while parsing the initial SQL.", err)
	}

	requests := strings.Split(string(file), ";")

	db := ConnectDB()
	defer util.DbClose(db)


	for _, request := range requests {
		if request != "" {
			_, err := db.Exec(request)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Database initialized.")

}

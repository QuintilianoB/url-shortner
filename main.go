package main

import (
	"github.com/subosito/gotenv"
	"log"
	"metro-ag/server"
	"metro-ag/store"
	"metro-ag/util"
	"os"
	"strconv"
)

func init() {

	// Load the .env file if it exists.
	if _, err := os.Stat(".env"); err == nil {
		err := gotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Set log configuration to show filenames and position in the log output.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	initMysql := util.GetEnv("dbInit", "false")
	t, err := strconv.ParseBool(initMysql)
	if err != nil {
		log.Fatal("Invalid configuration")
	}

	if t {
		// Load the initial sql.
		store.MysqlCreate()
	}

}

func main() {

	server.Server()

}

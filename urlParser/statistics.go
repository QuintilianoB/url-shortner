package urlParser

import (
	"log"
	"metro-ag/store"
	"metro-ag/util"
	"time"
)

type stat struct {
	Time int64 `json:"time"`
}

type Statistics struct {
	Url
	Data []stat `json:"data"`
}

// RegosterAccess will store the time of access for each recovered URL.
func (u Url) RegisterAccess() {

	var stat stat

	now := time.Now()
	stat.Time = now.Unix()

	db := store.ConnectDB()
	defer util.DbClose(db)

	stmt := "INSERT INTO statistics (id_url, time) VALUE (?, ?)"
	_, err := db.Query(stmt, u.Id, stat.Time)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Request for %s", u.Url)
}

/*
	Stats return an array with the time of each access to an URL.
	Possible returned error is:
		- DatabaseQueryError
 */
func (u Url) Stats() (Statistics, error) {

	var stats Statistics
	var stat stat

	// Append the Url info to the response.
	stats.Url.Url = u.Url
	u.SetPrefix()
	stats.UrlShort = u.UrlShort

	db := store.ConnectDB()
	defer util.DbClose(db)

	stmt := "SELECT time FROM statistics where id_url = ?"
	rows, err := db.Query(stmt, u.Id)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return stats, store.DatabaseQueryError
	}

	for rows.Next() {
		err = rows.Scan(&stat.Time)
		if err != nil {
			log.Println(err)
			return stats, store.DatabaseQueryError
		}
		stats.Data = append(stats.Data, stat)
	}

	// If there is and error at the rows recovery it will be logged but the statistics will be sent anyway.
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	log.Printf("Statistics recovered for %s", u.Url)
	return stats, nil
}

package urlParser

import (
	"database/sql"
	"encoding/json"
	"log"
	"metro-ag/store"
	"metro-ag/util"
	"net/http"
	"net/url"
)

// Errors for the Url type
var (
	InvalidJson     = util.ErrorMessage{"Invalid json."}
	InvalidUrl      = util.ErrorMessage{"Invalid url."}
	UrlNotFound     = util.ErrorMessage{"Url not found."}
	UrlAlreadyExist = util.ErrorMessage{"Url already exist."}
)

type Url struct {
	Id       int    `json:"-"`
	Url      string `json:"url"`
	UrlShort string `json:"url_short,omitempty"`
}

/*
	Validates verify if the json in the request is ok and if the received URL is a valid one.
	Possible returned errors are:
		- InvalidJson
		- InvalidUrl
*/
func (u *Url) Validate(r *http.Request) error {

	// Is there a problem with the json request?
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		log.Println(err)
		return InvalidJson
	}

	// Verifying if the URL is valid using the default lib of go.
	// https://stackoverflow.com/questions/25747580/ensure-a-uri-is-valid/25747925#25747925
	url, err := url.Parse(u.Url)
	if err != nil {
		log.Println(err)
		return InvalidUrl
	} else if !(url.Scheme != "" && url.Host != "") {
		return InvalidUrl
	}

	return nil
}

/*
	Get searches if the UrlShort or the URL  already exist on the database and populates the object with the result.
	Possible returned errors are:
		- UrlNotFound
        - DatabaseQueryError
*/
func (u *Url) Get() error {

	db := store.ConnectDB()
	defer util.DbClose(db)

	var stmt string
	var err error
	if u.UrlShort != "" {
		stmt = "SELECT * FROM urls where url_short = ?"
		err = db.QueryRow(stmt, u.UrlShort).Scan(&u.Id, &u.Url, &u.UrlShort)
	} else {
		stmt = "SELECT * FROM urls where url = ?"
		err = db.QueryRow(stmt, u.Url).Scan(&u.Id, &u.Url, &u.UrlShort)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return UrlNotFound
		}
		log.Println(err)
		return store.DatabaseQueryError
	}

	return nil
}

/*
	Create insert a new URL in the database.
	Possible returned errors are:
		- UrlAlreadyExist
        - DatabaseQueryError
*/
func (u *Url) Create() error {

	// Verify if the URL is already present. Returns the short URL if it does.
	err := u.Get()
	if err != UrlNotFound {
		if err == nil {
			log.Printf("Url %s already exist.", u.Url)
			return UrlAlreadyExist
		}
		return err
	}

	// Loop to grant that the new generated string does not exist.
	var tempUrl Url
	tempUrl.UrlShort = util.RandString(6)

Loop:
	for {
		err = tempUrl.Get()
		switch err {
		case UrlNotFound:
			u.UrlShort = tempUrl.UrlShort
			break Loop
		case nil:
			tempUrl.UrlShort = util.RandString(6)
		default:
			return err
		}
	}

	db := store.ConnectDB()
	defer util.DbClose(db)

	stmt := "INSERT INTO urls (url, url_short) value (?, ?)"
	_, err = db.Query(stmt, u.Url, u.UrlShort)

	if err != nil {
		log.Println(err)
		return store.DatabaseQueryError
	}

	log.Printf("Url %s saved as %s.", u.Url, u.UrlShort)

	return nil
}

// setPrefix append the url prefix to the saved url.
func (u *Url) SetPrefix() {
	suffix := util.GetEnv("frontendUrl", "http://127.0.0.1:8080")
	u.UrlShort = suffix + "/" + u.UrlShort
	log.Printf("Short Url %s", u.UrlShort)
}

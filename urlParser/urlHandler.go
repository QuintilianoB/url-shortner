package urlParser

import (
	"github.com/gorilla/mux"
	"metro-ag/store"
	"metro-ag/util"
	"net/http"
)

/*
	CreateUrl receives a request to shorting an Url and stores it.
	It expects a POST request in the route <ip>/create with a json containing:
		url: string containing the URL to be shortened.
	Returns 201 in case of success, 200 if the URL already exist in the database and a Json containing:
		url: string containing the original URL.
		url_short: string containing the URL in the shortened format.
*/
func CreateUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url Url

		err := url.Validate(r)
		if err != nil {
			util.SendError(w, http.StatusBadRequest, nil)
			return
		}

		err = url.Create()

		switch err {
		case UrlAlreadyExist:
			url.SetPrefix()
			util.SendSuccess(w, url)
		case store.DatabaseQueryError:
			util.SendError(w, http.StatusInternalServerError, err)
		case nil:
			// URL successfully shortened and stored.
			url.SetPrefix()
			w.WriteHeader(201)
			util.SendSuccess(w, url)
		default:
			util.SendError(w, http.StatusInternalServerError, nil)
		}
	}
}

/*
	AccessUrl parses a request for a short URL and process it.
	It expects a GET request in the route <ip>/{path} where path represents a short url created previously.
	In case of a valid path, returns a 200 and a JSON with to the original URL.
*/
func AccessUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url Url

		// Verifying if the short url exists in the database.
		vars := mux.Vars(r)
		url.UrlShort = vars["path"]

		err := url.Get()
		switch err {
		case UrlNotFound:
			util.SendError(w, http.StatusNotFound, nil)
		case store.DatabaseQueryError:
			util.SendError(w, http.StatusInternalServerError, err)
		case nil:
			url.RegisterAccess()
			url.UrlShort = ""
			util.SendSuccess(w, url)
			/*
				Which one is better?
				301 - We give a permanent redirection and the client does no consume resources on the server as often.
				302 - We ensure that every time that the client access the shortened URL we will register it.
			*/
			//http.Redirect(w, r, url.Url, 302)
		default:
			util.SendError(w, http.StatusInternalServerError, nil)
		}
	}
}

/*
	GetStats recovers the access statistics for a path.
	It expects a GET request in the route <ip>/{path}/stats where path represents a short url created previously.
	In case of a valid path, returns a 200 and a json in with following struct:
		{
			"url": "<ORIGINAL URL>",
    		"url_short": "<SHORTER URL>",
    		"data": [
        		{
            		"time": <UNIX TIMESTAMP>
        		},
        		...
				{
					"time": <UNIX TIMESTAMP>
				}
    		]
		}
*/
func GetStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url Url

		// Verifying if the short url exists in the database.
		vars := mux.Vars(r)
		url.UrlShort = vars["path"]

		err := url.Get()
		switch err {
		case UrlNotFound:
			util.SendError(w, http.StatusNotFound, nil)
		case nil:
			stats, err := url.Stats()
			if err == nil {
				util.SendSuccess(w, stats)
				break
			}
			fallthrough
		case store.DatabaseQueryError:
			util.SendError(w, http.StatusInternalServerError, err)
		default:
			util.SendError(w, http.StatusInternalServerError, nil)
		}
	}
}

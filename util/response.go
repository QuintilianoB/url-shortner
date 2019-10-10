package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendError(writer http.ResponseWriter, status int, error error) {
	writer.WriteHeader(status)
	writer.Header().Add("Content-type", "application/json")

	if error == nil {
		json.NewEncoder(writer)
	} else {
		err := json.NewEncoder(writer).Encode(error)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendSuccess(writer http.ResponseWriter, data interface{}) {
	writer.Header().Add("Content-type", "application/json")

	if data == nil {
		json.NewEncoder(writer)
	} else {
		err := json.NewEncoder(writer).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

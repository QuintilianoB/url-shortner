package server

import (
	"github.com/gorilla/mux"
	"metro-ag/urlParser"
)

func Routes(router *mux.Router) {

	router.HandleFunc("/create", urlParser.CreateUrl()).Methods("POST")
	router.HandleFunc("/{path}", urlParser.AccessUrl()).Methods("GET")
	router.HandleFunc("/{path}/stats", urlParser.GetStats()).Methods("GET")

}

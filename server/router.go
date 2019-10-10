package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"metro-ag/util"
	"log"
	"net/http"
)

func Server() {

	router := mux.NewRouter()

	Routes(router)

	host := util.GetEnv("serverAddr", "127.0.0.1")
	port := util.GetEnv("serverPort", "8000")
	serverUrl := fmt.Sprint(host + ":" + port)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowedMethods:         []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
	})

	fmt.Println("Start serving on", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, corsOpts.Handler(router)))
}

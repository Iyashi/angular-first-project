package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	developmentMode = true // set development mode to true

	// create the router and apply routes
	router := mux.NewRouter()
	router.Path("/hello").Methods(http.MethodGet).HandlerFunc(handleHello)
	router.Path("/work").Methods(http.MethodGet).HandlerFunc(handleWork)
	router.Path("/user/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(handleGetUser)
	router.Path("/users").Methods(http.MethodGet).HandlerFunc(handleGetUserList)

	// apply cors middleware to all routes to allow cross-origin requests from frontend to backend
	handler := CORSMiddleware(router)

	// start the http server
	listenAddr := fmt.Sprintf(":%v", backendPort)
	log.Printf("Listening on %s...\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, handler); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}

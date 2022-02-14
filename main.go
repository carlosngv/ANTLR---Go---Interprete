package main

import (
	"Parser/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	routes.UseRoutes(router)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	server := &http.Server{
		Handler:      corsWrapper.Handler(router),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}

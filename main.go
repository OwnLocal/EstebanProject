package main

import (
	"github.com/elohr/OwnLocal/config"
	"github.com/elohr/OwnLocal/router"
	esSetup "github.com/elohr/OwnLocal/setup"
	"log"
	"net/http"
)

// main is the entry point of the app
func main() {
	if config.DoSetup {
		log.Println("Will start ElasticSearch setup")

		err := esSetup.SetupES()

		if err != nil {
			log.Fatalf("Error setting up ElasticSearch Index: %s\n", err)
		}
	} else {
		// todo: start using TLS when authentication is implemented
		log.Printf("Starting server on %s", config.Config.WebURL)
		log.Fatal(http.ListenAndServe(config.Config.WebURL, router.Router))
	}
}

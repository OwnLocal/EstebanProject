package main

import (
	"log"
	"net/http"

	"github.com/OwnLocal/EstebanProject/config"
	"github.com/OwnLocal/EstebanProject/router"
	esSetup "github.com/OwnLocal/EstebanProject/setup"
)

// main is the entry point of the app
func main() {
	if config.DoSetup {
		log.Println("Will start ElasticSearch setup")

		err := esSetup.SetupES()

		if err != nil {
			log.Fatalf("Error setting up ElasticSearch Index: %s\n", err)
		}
	}

	// todo: start using TLS when authentication is implemented
	log.Printf("Starting server on %s", config.Config.WebURL)
	log.Fatal(http.ListenAndServe(config.Config.WebURL, router.Router))
}

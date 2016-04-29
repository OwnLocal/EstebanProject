package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type config struct {
	WebURL   string
	ES_URL   string
	ES_Index string
}

var Config config
var DoSetup bool

// init processes the flags provided when running this app, and based on that parses a config file into the config struct
func init() {
	prod := flag.Bool("p", false, "If the production flag is not specified, then the debug config file (config/config_debug.json) will be used.")
	flag.BoolVar(&DoSetup, "s", false, "When setup flag is passed the ElasticSearch index will be created. If it already exists it will be overwritten.")
	flag.Parse()

	var configFile []byte
	var err error

	if *prod {
		log.Println("Using production config file")
		configFile, err = ioutil.ReadFile("../../config/config_prod.json")
	} else {
		log.Println("Using debug config file")
		configFile, err = ioutil.ReadFile("../../config/config_debug.json")
	}

	if err != nil {
		log.Fatalf("Error parsing configuration file: %s\n", err)
	}

	json.Unmarshal(configFile, &Config)
}

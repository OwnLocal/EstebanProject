package setup

import (
	"encoding/csv"
	"github.com/elohr/OwnLocal/config"
	"github.com/elohr/OwnLocal/models/business"
	"gopkg.in/olivere/elastic.v3"
	"io"
	"log"
	"os"
)

// SetupES will setup de ES index (deleting it first if it already exists),
// and then will add all the business information into ES
func SetupES() error {
	client, err := elastic.NewSimpleClient(elastic.SetURL(config.Config.ES_URL))

	if err == nil {
		err = deleteIndexIfExists(client)

		if err == nil {
			// Create index
			err = createIndex(client)

			if err == nil {
				// Enter businesses information into ES
				err = processData()
			}
		}
	}

	return err
}

func deleteIndexIfExists(client *elastic.Client) error {
	// Check if index exist
	exists, err := client.IndexExists(config.Config.ES_Index).Do()

	// If it exists then delete it
	if err == nil && exists {
		if _, err = client.DeleteIndex(config.Config.ES_Index).Do(); err == nil {
			log.Println("Deleted ES index")
		}
	}

	// return the value of err (which can be nil)
	return err
}

func createIndex(client *elastic.Client) error {
	var err error

	if _, err = client.CreateIndex(config.Config.ES_Index).BodyJson(mapping).Do(); err == nil {
		log.Println("Created ES index")
	}

	return err
}

func processData() error {
	file, err := os.Open("./data/50k_businesses.csv")

	if err != nil {
		log.Fatalf("Error opening CSV file: %s\n", err)
	}

	reader := csv.NewReader(file)
	reader.Read() // to ignore the first line which are titles
	linesProcessed := 0

	log.Println("Started reading businesses file")

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error parsing line from CSV file: %s\n", err)
		}

		business, err := business.New(record...)

		if err != nil {
			log.Printf("Error creating new business: %s\n", err)
			continue
		}

		if err = business.Save(); err != nil {
			log.Println("Error saving business")
		}

		linesProcessed += 1
		if linesProcessed%500 == 0 {
			log.Printf("Lines processed: %d\n", linesProcessed)
		}
	}

	log.Printf("Finished processing CSV file: %d lines\n", linesProcessed)

	return err
}

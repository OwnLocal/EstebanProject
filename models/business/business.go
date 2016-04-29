package business

import (
	"log"
	"strconv"
	"time"

	"github.com/OwnLocal/EstebanProject/config"
	"gopkg.in/olivere/elastic.v3"
)

type business struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Address2  string    `json:"address2"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	Country   string    `json:"country"`
	Phone     string    `json:"phone"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	_ES_TYPE         = "business"
	_CSV_DATE_LAYOUT = "1/2/2006 15:04"
)

var searchFields = []string{"name", "address", "address2", "city", "state", "zip", "country", "phone"}

// New creates a new business object. All object properties need to be passed in the correct order,
// this simplifies creating a new business based on the input of the parsed CSV line
func New(params ...string) (*business, error) {
	b := new(business)

	id, err := strconv.Atoi(params[0]) // ignoring error
	if err != nil {
		log.Printf("Error parsing int for id: %s\n", params[0])
		return nil, err
	}

	b.Id = id
	b.Uuid = params[1]
	b.Name = params[2]
	b.Address = params[3]
	b.Address2 = params[4]
	b.City = params[5]
	b.State = params[6]
	b.Zip = params[7]
	b.Country = params[8]
	b.Phone = params[9]
	b.Website = params[10]

	createdAt, err := time.Parse(_CSV_DATE_LAYOUT, params[11]) // ignoring error
	if err != nil {
		log.Printf("Error parsing date for createdAt: %s\n", params[11])
		return nil, err
	}

	b.CreatedAt = createdAt

	return b, nil
}

// Search gets a search string, a from value that specifies how many results to skip,
// and max total number of results that can be returned.
func Search(search string, from int, size int) ([]interface{}, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(config.Config.ES_URL))

	if err != nil {
		return nil, err
	}

	var resp *elastic.SearchResult

	if search != "" {
		resp, err = client.Search().
			Index(config.Config.ES_Index).
			Query(elastic.NewMultiMatchQuery(search, searchFields...)).
			Sort("id", true).
			From(from).Size(size).
			Pretty(true).
			Do()
	} else {
		resp, err = client.Search().
			Index(config.Config.ES_Index).
			Sort("id", true).
			From(from).Size(size).
			Pretty(true).
			Do()
	}

	if err != nil {
		return nil, err
	}

	totalResults := len(resp.Hits.Hits)

	if totalResults == 0 {
		return nil, nil
	}

	results := make([]interface{}, totalResults)

	for i := 0; i < totalResults; i++ {
		results[i] = resp.Hits.Hits[i].Source
	}

	return results, nil
}

// GetAsJson returns an business and a bool value specifying if the business was found
// based on the id provided
func GetAsJson(id string) ([]byte, bool, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(config.Config.ES_URL))

	if err != nil {
		return nil, false, err
	}

	resp, err := client.Get().
		Index(config.Config.ES_Index).
		Type(_ES_TYPE).
		Id(id).
		Do()

	if err != nil {
		return nil, false, err
	}

	if !resp.Found {
		return nil, false, nil
	}

	respJson, err := resp.Source.MarshalJSON()

	return respJson, true, err
}

// Save indexes the business object in ElasticSearch
func (b business) Save() error {
	client, err := elastic.NewSimpleClient(elastic.SetURL(config.Config.ES_URL))

	if err != nil {
		return err
	}

	_, err = client.Index().
		Index(config.Config.ES_Index).
		Type(_ES_TYPE).
		Id(strconv.Itoa(b.Id)).
		BodyJson(b).
		Do()

	return err
}

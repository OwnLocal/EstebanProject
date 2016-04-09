package setup

var mapping = map[string]interface{}{
	"settings": map[string]interface{}{
		"index.number_of_shards":   1,
		"index.number_of_replicas": 0,
		"analysis": map[string]interface{}{
			"analyzer": map[string]interface{}{
				// Used to index name
				"nGram_Analyzer": map[string]interface{}{
					"type":      "custom",
					"tokenizer": "standard",
					"filter":    []string{"lowercase", "asciifolding", "nGram_filter"},
				},

				// Used for searches
				"whitespace_analyzer": map[string]interface{}{
					"type":      "custom",
					"tokenizer": "standard",
					"filter":    []string{"lowercase", "asciifolding", "length_filter"},
				},
			},
			"filter": map[string]interface{}{
				"nGram_filter": map[string]interface{}{
					"type":        "nGram",
					"min_gram":    1,
					"max_gram":    20,
					"token_chars": []string{"letter", "digit"},
				},
				"length_filter": map[string]interface{}{
					"type": "length",
					"min":  1,
					"max":  20,
				},
			},
		},
	},

	"mappings": map[string]interface{}{
		"business": map[string]interface{}{
			"dynamic": false,
			"_all": map[string]interface{}{
				"enabled": false,
			},
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type": "integer",
				},
				"uuid": map[string]interface{}{
					"type":  "string",
					"index": "no",
				},
				"name": map[string]interface{}{
					"type":            "string",
					"analyzer":        "nGram_Analyzer",
					"search_analyzer": "whitespace_analyzer",
				},
				"address": map[string]interface{}{
					"type": "string",
				},
				"address2": map[string]interface{}{
					"type": "string",
				},
				"city": map[string]interface{}{
					"type": "string",
				},
				"state": map[string]interface{}{
					"type": "string",
				},
				"zip": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"country": map[string]interface{}{
					"type": "string",
				},
				"phone": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"website": map[string]interface{}{
					"type":  "string",
					"index": "no",
				},
				"created_at": map[string]interface{}{
					"type":  "date",
					"index": "no",
				},
			},
		},
	},
}

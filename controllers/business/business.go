package business

import (
	//"encoding/json"
	"log"
	"net/http"
	//"strconv"

	"github.com/OwnLocal/EstebanProject/models/business"
	"github.com/OwnLocal/EstebanProject/util"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
)

const (
	_RESULT_SIZE = 50
)

// List send the client the results of the businesses search.
// It takes three optional parameters: a search string,
// a from value indicating how many results to skip (paging),
// and a size value that specified that max number of results to include
// Metadata included on the response: the search text, the from value used,
// the maxResults (size) used, the number of results being returned,
// and an array with the businesses that are part of the result
//func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	searchText := p.ByName("search")
//	from, _ := strconv.Atoi(p.ByName("from")) // returns 0 if non numeric value is given
//	size, _ := strconv.Atoi(p.ByName("size")) // returns 0 if non numeric value is given
//
//	if size == 0 {
//		size = _RESULT_SIZE
//	}
//
//	businesses, err := business.Search(searchText, from, size)
//
//	if err != nil {
//		log.Printf("Error on business.List: %s\n", err)
//		util.WriteServerError(w)
//		return
//	}
//
//	response := map[string]interface{}{
//		"search":     searchText,
//		"from":       from,
//		"maxResults": size,
//		"results":    len(businesses),
//		"businesses": businesses,
//	}
//
//	responseJson, err := json.Marshal(response)
//
//	if err != nil {
//		log.Printf("Error on business.List: %s\n", err)
//		util.WriteServerError(w)
//		return
//	}
//
//	util.WriteSuccess(w, responseJson)
//}

func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	from, _ :=  strconv.Atoi(r.URL.Query().Get("from"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	if size == 0 {
		size = _RESULT_SIZE
	}

	res := fmt.Sprintf("%v %v", from, size)

	util.WriteSuccess(w, []byte(res))
}

// Get returns the business object based on the id provided
func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	businessJson, found, err := business.GetAsJson(p.ByName("id"))

	if err != nil {
		log.Printf("Error on business.Get: %s\n", err)
		util.WriteServerError(w)
		return
	}

	if !found {
		util.WriteNotFound(w)
		return
	}

	util.WriteSuccess(w, businessJson)
}

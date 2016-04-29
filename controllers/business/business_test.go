package business_test

import (
	"testing"
	"net/http"
	"log"
	"net/http/httptest"
	"github.com/OwnLocal/EstebanProject/controllers/business"
)


func TestQueryString(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888/businesses?from=0&size=25", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := business.List
	handler(w, req, nil)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestQueryStringDefaults(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888/businesses", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := business.List
	handler(w, req, nil)

	if w.Body.String() != "0 50" {
		t.Fail()
	}
}

func TestQueryStringValues(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888/businesses?from=0&size=25", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := business.List
	handler(w, req, nil)

	resBody := w.Body.String()
	t.Log("Error: ", resBody)

	if resBody != "0 25" {
		t.Fail()
	}
}
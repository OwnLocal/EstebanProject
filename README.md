## OwnLocal
## Demo!
There is a demo running on http://ownlocal.r6.io:30205.

###Endpoints

**Search**

`GET /businesses` returns a search result of the businesses. This endpoint may take several parameters: a search string, a number of items to skip (for paging), and the max number of results per page. Apart from the call with no parameters, the endpoint can be called like this `GET /businesses/:searchString/:from/:size`. Any of these parameters can be blank. `:from` by default is 0, that means it is the first page, and `:size` by default is 50 results. Results are always ordered by `id`.

Searches are made on most fields at the same time. Name of the business has an `ngram analyzer`, which means searches can match partial words, other fields use `standard analyzer` which means only full words will be matched when searched.

_Search Examples:_

`GET /businesses` will return the first 50 businesses

`GET /businesses/bot` will return the businesses that match the search `bot`.

`GET /businesses/bot/10/5` will skip 10 results, and return 5 results (or page 3 if each page includes 5 results) that match the search `bot`.

`GET /businesses//10/5` will skip 10 results, and return 5 results (or page 3 if each page includes 5 results) of all businesses.

`GET /businesses/6068112791` will return the businesses with phone number `6068112791`.

`GET /businesses/Roobland` will return the businesses with city `Roobland`.

Results include an object with the following structure:

```javascript
{
  businesses: [...],  // array with results
  from: 0,            // number of results being skipped
  maxResults: 50,     // max size of page
  results: 12,        // the actual results of the page
  search: bot         // the search string that generated these results   
}
```

**Get by ID**

`GET /business/:id` will return the business with that has the provided `id`.

_Example:_

`GET /business/33` will return the business with id `33`, which is Mraz and Sons.


### Stack
Go and ElasticSearch

**Libraries:**

elastic.v3 (https://github.com/olivere/elastic/tree/v3.0.30)

httprouter (https://github.com/julienschmidt/httprouter)

### Dependencies
**Go**

Go is required to run the project. For instructions on how to install: https://golang.org/doc/install

**ElasticSearch**

ElasticSearch is also required. https://www.elastic.co/downloads/elasticsearch

**Libraries**

`go get github.com/julienschmidt/httprouter`

`go get gopkg.in/olivere/elastic.v3`

### Installation

Clone the project from git. The project has alrady been built, but there are some config files that may need be be modified.

**API and ElasticSearch running on same local machine**

On `config/config_debug.json` change the value of ES_URL to either `http://localhost:9200`, or to match the IP where ElasticSearch is listening.

**Deploy to Production**

If service will be deployed to another server then change the values of `config/config_prod.json` accordingly.

### Run

Execute `./OwnLocal` with the required flags. The available flags are:

`-p` to use prod config file. If this flag is not used then the debug config file will be used by default.

`-s` to setup ES. This creates the ES index (deleting it first if it already exists), adds the index settings and mappings, and then parses the CSV file to index all businesses into ES.

### Project Structure
**Config:** containes configuration files to connect to ES and to define listening IP/port of API.

**Controllers:** controller files for API requests.

**Data:** source CSV file.

**Models:** description for business entity.

**Router:** handler for URL routes.

**Setup:** files to setup ES and index data.

**Util:** other files, for now it only has wrapper for ResponseWriter.

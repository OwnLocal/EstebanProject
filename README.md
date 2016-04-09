## OwnLocal
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

Clone the project from git. The Project has alrady been built, but there are some config files that may need be be modified.

**API and ElasticSearch running on same local machine**

On `config/config_debug.json` change the value of ES_URL to either `http://localhost:9200`, or to match the IP where ElasticSearch is listening.

**Deploy to Production**

If service will be deployed to another server then change the values of `config/config_prod.json` accordingly.

### Run

Execute `./OwnLocal` with the required flags. The available flags are:

`-p` to use prod config file.
`-s` to setup ES.

### Project Structure
**Config:** containes configuration files to connect to ES and to define listening IP/port of API.

**Controllers:** controller files for API requests.

**Data:** source CSV file.

**Models:** description for business entity.

**Router:** handler for URL routes.

**Setup:** files to setup ES and index data.

**Util:** other files, for now it only has wrapper for ResponseWriter.



// elasticsearch/client.go

package elasticsearch

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

var Client *elasticsearch.Client

func Initialize() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	var err error
	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed to creating elasticsearch client: %s", err)
	}
}

package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func Search(query map[string]interface{}, index string) (*esapi.Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	return Client.Search(
		Client.Search.WithContext(context.Background()),
		Client.Search.WithIndex(index),
		Client.Search.WithBody(&buf),
		Client.Search.WithTrackTotalHits(true),
	)
}

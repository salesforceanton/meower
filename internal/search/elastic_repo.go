package search_repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/salesforceanton/meower/internal/schema"
)

const (
	MEOWER_INDEX_NAME = "meower"
	MEOWS_TABLE_NAME  = "meows"
)

type ElasticRepo struct {
	client *elasticsearch.Client
}

func NewElasticRepo(addr string) (*ElasticRepo, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s", addr)},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ElasticRepo{client: client}, nil
}

func (r *ElasticRepo) Close() {
	r.client.Indices.Close([]string{MEOWER_INDEX_NAME})
}

func (r *ElasticRepo) InsertMeow(ctx context.Context, message schema.Meow) error {
	successChan := make(chan int)
	errorChan := make(chan error)

	go func() {
		data, err := json.Marshal(message)
		if err != nil {
			errorChan <- err
		}
		_, err = r.client.Index(MEOWER_INDEX_NAME, bytes.NewReader(data))
		if err != nil {
			errorChan <- err
		} else {
			successChan <- 1
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errorChan:
		return err
	case <-successChan:
		return nil
	}
}

func (r *ElasticRepo) SearchMeows(ctx context.Context, queryString string, skip, take int64) ([]schema.Meow, error) {
	successChan := make(chan []schema.Meow)
	errorChan := make(chan error)

	go func() {
		// Build query
		query := fmt.Sprintf(
			`{
				"from": %d,
	  			"size": %d,
				"query": {
					"multi_match": {
					  	"query": "%s",
					  	"fields": [
							"body",
							"created_at"
					  	]
					}
				}
			}`, skip, take, queryString)

		// Make request to Elastic
		res, err := r.client.Search(
			r.client.Search.WithIndex(MEOWER_INDEX_NAME),
			r.client.Search.WithBody(strings.NewReader(query)),
		)
		if err != nil {
			errorChan <- err
			return
		}

		// Parse result
		var jsonResponse map[string]interface{}
		var searchResult []schema.Meow
		defer res.Body.Close()
		if err = json.NewDecoder(res.Body).Decode(&jsonResponse); err != nil {
			errorChan <- err
		} else {
			for _, hit := range jsonResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
				source := hit.(map[string]interface{})["_source"].(map[string]interface{})
				createdAt, _ := time.Parse(time.RFC3339, source["createdAt"].(string))

				searchResult = append(searchResult, schema.Meow{
					Id:        source["id"].(string),
					Body:      source["body"].(string),
					CreatedAt: createdAt,
				})
			}
			successChan <- searchResult
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errorChan:
		return nil, err
	case result := <-successChan:
		return result, nil
	}
}

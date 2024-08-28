package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"review-job/internal/conf"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

// ReviewWorker implements transport.Server
type ReviewWorker struct {
	kafkaReader *kafka.Reader
	esClient    *ElasticSearchClient
	log         *log.Helper
}

type ElasticSearchClient struct {
	client *elasticsearch.TypedClient
	Index  string
}

type KafkaMessage struct {
	Type     string                   `json:"type"`
	Database string                   `json:"database"`
	Table    string                   `json:"table"`
	IsDdl    bool                     `json:"isDdl"`
	Data     []map[string]interface{} `json:"data"`
}

func (w *ReviewWorker) Start(ctx context.Context) error {
	for {
		m, err := w.kafkaReader.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		}

		km := new(KafkaMessage)
		if err := json.Unmarshal(m.Value, km); err != nil {
			continue
		}

		if km.Type == "INSERT" {
			w.log.Infof("Insert data: %v", km.Data)
			for _, data := range km.Data {
				if err := w.esClient.IndexDocument(data); err != nil {
					w.log.Errorf("Index document failed: %v", err)
				}
			}
		} else {
			w.log.Infof("Update data: %v", km.Data)
			for _, data := range km.Data {
				if err := w.esClient.UpdateDocument(data); err != nil {
					w.log.Errorf("Update document failed: %v", err)
				}
			}
		}
	}
}

func (w *ReviewWorker) Stop(ctx context.Context) error {
	w.kafkaReader.Close()
	return nil
}

func (es *ElasticSearchClient) IndexDocument(data map[string]interface{}) error {
	reviewID := data["review_id"].(string)
	resp, err := es.client.Index(es.Index).Id(reviewID).Document(data).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(resp.Result)
	return nil
}

func (es *ElasticSearchClient) UpdateDocument(data map[string]interface{}) error {
	reviewID := data["review_id"].(string)
	resp, err := es.client.Update(es.Index, reviewID).Doc(data).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(resp.Result)
	return nil
}

func NewReviewWorker(kafkaReader *kafka.Reader, esClient *ElasticSearchClient, logger log.Logger) *ReviewWorker {
	return &ReviewWorker{
		kafkaReader: kafkaReader,
		esClient:    esClient,
		log:         log.NewHelper(logger),
	}
}

func NewKafkaReader(c *conf.Kafka) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.Brokers,
		GroupID: c.GroupId,
		Topic:   c.Topic,
	})
}

func NewElasticSearchClient(c *conf.ElasticSearch) (*ElasticSearchClient, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: c.Addrs,
	})
	if err != nil {
		return nil, err
	}
	return &ElasticSearchClient{
		client: client,
		Index:  c.Index,
	}, nil
}

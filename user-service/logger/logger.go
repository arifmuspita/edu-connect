package logger

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type ElasticHook struct {
	client *elastic.Client
	index  string
}

func NewElasticHook(client *elastic.Client, index string) *ElasticHook {
	return &ElasticHook{
		client: client,
		index:  index,
	}
}

func (hook *ElasticHook) Fire(entry *logrus.Entry) error {

	doc := make(map[string]interface{})
	for k, v := range entry.Data {
		doc[k] = v
	}

	doc["message"] = entry.Message
	doc["level"] = entry.Level.String()
	doc["@timestamp"] = time.Now()

	_, err := hook.client.Index().
		Index(hook.index).
		BodyJson(doc).
		Do(context.Background())

	return err
}

func (hook *ElasticHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func InitLogger() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.JSONFormatter{})

	elasticURL := os.Getenv("ESHOST")
	elasticUser := os.Getenv("ESUSER")
	elasticPass := os.Getenv("ESPASS")

	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetBasicAuth(elasticUser, elasticPass),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Elasticsearch: %v", err)
	}

	hook := NewElasticHook(client, "user-service-logs")
	Log.Hooks.Add(hook)

	Log.Info("Logger initialized with Elasticsearch hook")
}

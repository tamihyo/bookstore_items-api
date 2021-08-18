package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
	"github.com/tamihyo/bookstore_utils-go/logger"
)

//when creating variable, we need to utilize the client attributes
var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

//making struct based on return of package function
type esClient struct {
	client *elastic.Client //the actual library that we implemented
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}
	//client is of type ES client interface
	//setting the client by putting our private method on the interface
	Client.setClient(client)

	//create index if not exists
}

/*
working with the actual object and not copy of the object
but c.client is poiner of elastic     3
*/
func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

//implement index function
func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index("items").
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in %s", index), err)
		return nil, err
	}
	return result, nil
}

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
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, docType string, id string) (*elastic.GetResult, error)
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
func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		fmt.Println(result.Found)
		logger.Error(fmt.Sprintf("error wheren trying to get id %s", id), err)
		return nil, err
	}

	//doc and err not exists
	if !result.Found {
		return nil, nil
	}

	return result, nil

}

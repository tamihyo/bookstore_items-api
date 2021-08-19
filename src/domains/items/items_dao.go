package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tamihyo/bookstore_items-api/src/clients/elasticsearch"
	"github.com/tamihyo/bookstore_items-api/src/domains/queries"
	"github.com/tamihyo/bookstore_utils-go/rest_errors"
)

//save whatever the items without  no matter the db using is
//if wanna change the database, change item_dto only

const (
	typeItem   = "item"
	indexItems = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewInternalServerError(fmt.Sprintf("no item found with  id %s", i.Id), errors.New("database error"))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with the id %s ", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search docs", errors.New("database error"))
	}
	fmt.Println(result)

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error wheren trying to parse response", errors.New("database error"))
		}
		items[index].Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no itmes foundmatchin given criteria")
	}
	return items, nil
}

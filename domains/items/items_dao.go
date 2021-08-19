package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tamihyo/bookstore_items-api/clients/elasticsearch"
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

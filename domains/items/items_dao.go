package items

import (
	"errors"

	"github.com/tamihyo/bookstore_items-api/clients/elasticsearch"
	"github.com/tamihyo/bookstore_utils-go/rest_errors"
)

//save whatever the items without  no matter the db using is
//if wanna change the database, change item_dto only

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, 1)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.Id = result.Id
	return nil
}

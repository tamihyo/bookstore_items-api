package services

import (
	"github.com/tamihyo/bookstore_items-api/src/domains/items"
	"github.com/tamihyo/bookstore_items-api/src/domains/queries"
	"github.com/tamihyo/bookstore_utils-go/bookstore_utils-go/rest_errors"
)

/*
create variable calles item service being a type of item
service interface
the value of ItemsService is the pointer of itemsService being struct
*/
var (
	ItemsService itemsServiceInterface = &itemsService{} //default type value is nil
)

//basic interface that we are going to develop
type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
	// return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	// dao.Search(query)
	return dao.Search(query)
}

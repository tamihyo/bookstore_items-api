package services

import (
	"net/http"

	"github.com/tamihyo/bookstore_items-api/domains/items"
	"github.com/tamihyo/bookstore_utils-go/rest_errors"
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
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError{
		"implement me",http.StatusNotImplemented,"not_implemented",nil
	}
}

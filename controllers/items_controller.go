package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tamihyo/bookstore_items-api/domains/items"
	"github.com/tamihyo/bookstore_items-api/services"
	"github.com/tamihyo/bookstore_items-api/utils/http_utils"
	"github.com/tamihyo/bookstore_users-go/oauth"
	"github.com/tamihyo/bookstore_utils-go/rest_errors"
)

var (
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

//defining as method
type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//todo : return error to the caller
		http_utils.RespondJson(w, err.Status, *err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}

	defer r.Body.Close() // close body requestS

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	//todo return created items as json with http status 201 -created
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)

}

package app

import (
	"net/http"

	"github.com/tamihyo/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)

}

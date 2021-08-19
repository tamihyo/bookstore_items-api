package application

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tamihyo/bookstore_items-api/src/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
	elasticsearch.Init()
	//create server
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8085",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Microsecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

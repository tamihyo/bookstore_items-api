package main

import (
	"os"

	"github.com/tamihyo/bookstore_items-api/app"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}

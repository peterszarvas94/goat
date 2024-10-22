package main

import (
	"fmt"

	"github.com/peterszarvas94/goat/modules/handlers"
	"github.com/peterszarvas94/goat/modules/routing"
	"github.com/peterszarvas94/goat/templates/pages"
)

func main() {
	router := routing.NewRouter()

	router.GetTempl("/{$}", pages.Index())
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	err := router.Serve("localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

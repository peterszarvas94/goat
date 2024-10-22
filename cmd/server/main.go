package main

import (
	"fmt"
	"strings"

	"github.com/peterszarvas94/goat/config"
	"github.com/peterszarvas94/goat/modules/handlers"
	"github.com/peterszarvas94/goat/modules/routing"
	"github.com/peterszarvas94/goat/templates/pages"
)

func main() {
	router := routing.NewRouter()

	router.GetTempl("/{$}", pages.Index())
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	url := strings.Join([]string{"localhost", config.PORT}, ":")

	err := router.Serve(url)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

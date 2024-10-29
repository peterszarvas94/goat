package main

import (
	"fmt"
	"project/config"
	"project/handlers"
	"project/templates/pages"
	"strings"

	"github.com/peterszarvas94/goat/logging"
	"github.com/peterszarvas94/goat/routing"
)

func main() {
	logging.Logger.Debug("hello")
	router := routing.NewRouter()

	router.GetTempl("/{$}", pages.Index())
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	url := strings.Join([]string{"localhost", config.Port}, ":")

	err := router.Serve(url)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

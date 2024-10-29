package main

import (
	"fmt"
	"project/handlers"
	"project/templates/pages"
	"strings"

	c "github.com/peterszarvas94/goat/config"
	"github.com/peterszarvas94/goat/routing"
)

func main() {
	router := routing.NewRouter()

	router.GetTempl("/{$}", pages.Index())
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	config := c.NewConfig()

	url := strings.Join([]string{"localhost", config.Port}, ":")

	err := router.Serve(url)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

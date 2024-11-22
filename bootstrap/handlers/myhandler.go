package handlers

import (
	"fmt"
	"net/http"

	"bootstrap/templates/pages"

	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/server"
)

func MyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := database.Get()
	if err != nil {
		server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "This is a response from http.HandlerFunc!")
}

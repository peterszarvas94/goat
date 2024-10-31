package handlers

import (
	"fmt"
	"net/http"
	"project/templates/pages"

	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/router"
)

func MyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := database.Connect()
	if err != nil {
		logger.Error("can not connect to DB")
		router.ShowTempl(pages.ServerError(), w, r, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "This is a response from http.HandlerFunc!")
}

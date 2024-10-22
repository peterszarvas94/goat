package handlers

import (
	"fmt"
	"net/http"
)

func MyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a response from http.HandlerFunc!")
}

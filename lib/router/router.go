package router

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

type routerT struct {
	mux *http.ServeMux
}

var r = &routerT{
	mux: http.NewServeMux(),
}

func addRoute(method string, path string, handler http.Handler) {
	pattern := strings.Join([]string{method, path}, " ")
	r.mux.Handle(pattern, handler)
}

func Get(path string, handler http.HandlerFunc) {
	addRoute("GET", path, handler)
}
func Post(path string, handler http.HandlerFunc) {
	addRoute("POST", path, handler)
}
func Patch(path string, handler http.HandlerFunc) {
	addRoute("PATCH", path, handler)
}
func Delete(path string, handler http.HandlerFunc) {
	addRoute("DELETE", path, handler)
}

func Templ(path string, component templ.Component) {
	addRoute("GET", path, templ.Handler(component))
}

func Serve(addr string) error {
	if err := http.ListenAndServe(addr, r.mux); err != nil {
		return err
	}
	return nil
}

func ShowTempl(component templ.Component, w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	templ.Handler(component).ServeHTTP(w, r)
}

package routing

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

type router struct {
	mux *http.ServeMux
}

func NewRouter() *router {
	mux := http.NewServeMux()

	return &router{
		mux: mux,
	}
}

var Router = NewRouter()

func (r *router) addRoute(method string, path string, handler http.Handler) {
	pattern := strings.Join([]string{method, path}, " ")
	r.mux.Handle(pattern, handler)
}

func (r *router) Get(path string, handler http.HandlerFunc) {
	r.addRoute("GET", path, handler)
}
func (r *router) Post(path string, handler http.HandlerFunc) {
	r.addRoute("POST", path, handler)
}
func (r *router) Patch(path string, handler http.HandlerFunc) {
	r.addRoute("PATCH", path, handler)
}
func (r *router) Delete(path string, handler http.HandlerFunc) {
	r.addRoute("DELETE", path, handler)
}

func (r *router) GetTempl(path string, component templ.Component) {
	r.addRoute("GET", path, templ.Handler(component))
}

func (r *router) Serve(addr string) error {
	if err := http.ListenAndServe(addr, r.mux); err != nil {
		return err
	}
	return nil
}

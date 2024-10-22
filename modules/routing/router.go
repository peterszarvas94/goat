package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/goat/templates/pages"
)

type router struct {
	mux *http.ServeMux
}

func NewRouter() *router {
	mux := http.NewServeMux()

	router := &router{
		mux: mux,
	}

	router.GetTempl("/", pages.NotFound())

	return router
}

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
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(addr, r.mux); err != nil {
		return err
	}
	return nil
}

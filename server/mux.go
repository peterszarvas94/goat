package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	l "github.com/peterszarvas94/goat/logger"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	Mux         http.ServeMux
	middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{
		Mux:         *http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) applyMiddlewares(handler http.HandlerFunc) http.HandlerFunc {
	for i := len(r.middlewares) - 1; i >= 0; i-- { // Apply in reverse order
		handler = r.middlewares[i](handler)
	}
	return handler
}

func (r *Router) addRoute(method string, path string, handler http.HandlerFunc) {
	pattern := strings.Join([]string{method, path}, " ")
	wrappedHandler := r.applyMiddlewares(handler)
	r.Mux.Handle(pattern, wrappedHandler)
	l.Logger.Debug("Route added", slog.String("method", method), slog.String("path", path))
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.addRoute("GET", path, handler)
}
func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.addRoute("POST", path, handler)
}
func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.addRoute("PATCH", path, handler)
}
func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.addRoute("DELETE", path, handler)
}

func (r *Router) TemplGet(path string, component templ.Component) {
	r.addRoute("GET", path, templ.Handler(component).ServeHTTP)
}

func (r *Router) Favicon(filePath string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
	r.addRoute("GET", "/favicon.ico", handler)
}

func Render(w http.ResponseWriter, r *http.Request, component templ.Component, status int) {
	w.WriteHeader(status)
	component.Render(r.Context(), w)
}

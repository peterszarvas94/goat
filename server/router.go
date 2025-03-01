package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/goat/logger"
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

func (r *Router) Use(mws ...Middleware) {
	r.middlewares = append(r.middlewares, mws...)
}

func (r *Router) applyMiddlewares(handler http.HandlerFunc, routeMiddlewares ...Middleware) http.HandlerFunc {
	// Apply route-specific middlewares first
	for i := len(routeMiddlewares) - 1; i >= 0; i-- {
		handler = routeMiddlewares[i](handler)
	}

	// Apply global middlewares
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	return handler
}

func (r *Router) addRoute(method string, path string, handler http.HandlerFunc, routeMiddlewares ...Middleware) {
	pattern := strings.Join([]string{method, path}, " ")
	wrappedHandler := r.applyMiddlewares(handler, routeMiddlewares...)
	r.Mux.Handle(pattern, wrappedHandler)
	logger.Debug("Route added", slog.String("method", method), slog.String("path", path))
}

func (r *Router) Get(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute("GET", path, handler, middlewares...)
}

func (r *Router) Post(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute("POST", path, handler, middlewares...)
}

func (r *Router) Patch(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute("PATCH", path, handler, middlewares...)
}

func (r *Router) Delete(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute("DELETE", path, handler, middlewares...)
}

func (r *Router) TemplGet(path string, component templ.Component, middlewares ...Middleware) {
	handler := templ.Handler(component).ServeHTTP
	r.addRoute("GET", path, handler, middlewares...)
}

func (r *Router) Favicon(filePath string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
	r.addRoute("GET", "/favicon.ico", handler)
}

func (r *Router) Static(route, folder string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix(route, http.FileServer(http.Dir(folder)))
		fs.ServeHTTP(w, r)
	}
	r.addRoute("GET", route, handler)
}

func Render(w http.ResponseWriter, r *http.Request, component templ.Component, status int) {
	w.WriteHeader(status)
	component.Render(r.Context(), w)
}

// Set up the following static routes
//
// favicon.ico
//
// scripts folder
//
// styles folder
func (r *Router) Setup() {
	r.Favicon("favicon.ico")
	r.Static("/scripts/", "./scripts")
	r.Static("/styles/", "./styles")
}

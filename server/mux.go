package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	l "github.com/peterszarvas94/goat/logger"
)

type Mux struct {
	mux *http.ServeMux
	url string
}

func NewMux(url string) *Mux {
	return &Mux{
		url: url,
		mux: http.NewServeMux(),
	}
}

func (m *Mux) addRoute(method string, path string, handler http.Handler) {
	pattern := strings.Join([]string{method, path}, " ")
	m.mux.Handle(pattern, handler)
	l.Logger.Debug("Route added", slog.String("method", method), slog.String("path", path))
}

func (m *Mux) Get(path string, handler http.HandlerFunc) {
	m.addRoute("GET", path, handler)
}
func (m *Mux) Post(path string, handler http.HandlerFunc) {
	m.addRoute("POST", path, handler)
}
func (m *Mux) Patch(path string, handler http.HandlerFunc) {
	m.addRoute("PATCH", path, handler)
}
func (m *Mux) Delete(path string, handler http.HandlerFunc) {
	m.addRoute("DELETE", path, handler)
}

func (m *Mux) TemplGet(path string, component templ.Component) {
	m.addRoute("GET", path, templ.Handler(component))
}

func Render(w http.ResponseWriter, r *http.Request, component templ.Component, status int) {
	w.WriteHeader(status)
	component.Render(r.Context(), w)
}

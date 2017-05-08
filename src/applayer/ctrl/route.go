package ctrl

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Handle struct {
	Pattern string
	Handler func(http.ResponseWriter, *http.Request)
	Methods []string
}

type Route struct {
	pathPrefix  string
	StrictSlash bool
	Handles     []Handle
	Middlewares []negroni.Handler
}

// NewRoute returns a pointer of Route.
func NewRoute(pathPrefix string) *Route {
	return &Route{
		pathPrefix: pathPrefix,
	}
}

// PathPrefix returns the r.PathPrefix.
func (r Route) PathPrefix() string {
	return r.pathPrefix
}

// Handler returns a http.Handler.
func (r Route) Handler() http.Handler {
	router := mux.NewRouter().PathPrefix(r.PathPrefix()).Subrouter().StrictSlash(r.StrictSlash)

	for _, handle := range r.Handles {
		// sets handler func.
		router.HandleFunc(handle.Pattern, handle.Handler).Methods(handle.Methods...)
	}

	// sets middlewares.
	return negroni.New(
		append(r.Middlewares, negroni.Wrap(router))...,
	)
}

package ctrl

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Muxer interface {
	PathPrefix() string
	Handler() http.Handler
}

// NewHandler returns a http.Handler by the given list.
func NewHandler(mx ...Muxer) http.Handler {
	router := mux.NewRouter()

	for _, m := range mx {
		router.PathPrefix(m.PathPrefix()).Handler(m.Handler())
	}

	return router
}

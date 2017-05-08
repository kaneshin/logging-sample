package ctrlapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/kaneshin/logging-sample/src/app/ctrl"
	"github.com/kaneshin/logging-sample/src/track"
)

var (
	V1Route *ctrl.Route
	render  = ctrl.Render
)

func init() {
	V1Route = ctrl.NewRoute("/1.0")
	V1Route.StrictSlash = true
	V1Route.Middlewares = []negroni.Handler{
		negroni.NewLogger(),
		negroni.HandlerFunc(ctrl.AuthMiddleware),
		negroni.HandlerFunc(ctrl.EventMiddleware),
	}
}

func init() {
	V1Route.Handles = []ctrl.Handle{
		{
			Pattern: "/log",
			Handler: serveLog,
			Methods: []string{http.MethodPost},
		},
	}
}

func serveLog(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Level   string `json:"level"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error_message": err.Error()})
		return
	}

	render.JSON(w, http.StatusCreated, map[string]string{"status": "ok"})

	go func() {
		track.AppLog(track.AppData{
			Level:   p.Level,
			Message: p.Message,
			Time:    time.Now(),
		})
	}()
}

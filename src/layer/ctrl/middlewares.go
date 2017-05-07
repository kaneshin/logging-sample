package ctrl

import (
	"net/http"
	"strings"
	"time"

	"github.com/kaneshin/logging-sample/src/track"
)

const thatToken = "Bearer thisistoken"

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("Authorization")

	switch {
	case token == "":
		rw.Header().Set("WWW-Authenticate", "Bearer realm=\"no_token\"")
		Render.JSON(rw, http.StatusUnauthorized, map[string]string{"status": "unauthorized"})
		return

	case token != thatToken:
		rw.Header().Set("WWW-Authenticate", "Bearer error=\"invalid_request\"")
		Render.JSON(rw, http.StatusBadRequest, map[string]string{"status": "bad request"})
		return
	}

	next(rw, r)

	go func() {
		track.EventLog(track.EventData{
			Category: "Auth",
			Action:   "Login",
			Value:    strings.Fields(token)[1],
			Time:     time.Now(),
		})
	}()
}

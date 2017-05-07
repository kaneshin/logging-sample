package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kaneshin/logging-sample/src/layer/ctrl"
	"github.com/kaneshin/logging-sample/src/layer/ctrl/ctrladmin"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	handler := ctrl.NewHandler(
		ctrladmin.Route,
	)

	log.Fatal(http.ListenAndServe(*addr, handler))
}

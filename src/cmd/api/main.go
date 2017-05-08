package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kaneshin/logging-sample/src/applayer/ctrl"
	"github.com/kaneshin/logging-sample/src/applayer/ctrl/ctrlapi"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	handler := ctrl.NewHandler(
		ctrlapi.V1Route,
	)

	log.Fatal(http.ListenAndServe(*addr, handler))
}

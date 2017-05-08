package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/kaneshin/logging-sample/src/track"
)

func main() {
	st := flag.String("struct", "", "(app|event)")
	flag.Parse()

	switch *st {
	case "app":
		printSchema(track.AppData{})
	case "event":
		printSchema(track.EventData{})
	default:
		fmt.Fprintln(os.Stderr, "[ERROR] Use struct option")
	}
}

func printSchema(st interface{}) {
	is, err := bigquery.InferSchema(st)
	if err != nil {
		panic(err)
	}

	var list []interface{}
	for _, fs := range is {
		t := struct {
			Name string `json:"name"`
			Type string `json:"type"`
		}{
			strings.ToLower(fs.Name),
			string(fs.Type),
		}
		list = append(list, t)
	}

	b, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, string(b))
}

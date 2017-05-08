package bq

//go:generate sh -c "go run ./../../src/cli/bqschema/main.go -struct app > ./schema/app.json"
//go:generate sh -c "go run ./../../src/cli/bqschema/main.go -struct event > ./schema/event.json"

package track

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

const AppLogFilename = "/tmp/logging-sample.app.log"

var appLog *log.Logger

func init() {
	f, err := os.OpenFile(AppLogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	appLog = log.New(f, "", 0)
}

type AppData struct {
	Level   string    `json:"level"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (d AppData) String() string {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(d); err != nil {
		return err.Error()
	}
	return buf.String()
}

func AppLog(d AppData) {
	appLog.Print(d.String())
}

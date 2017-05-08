package track

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

const EventLogFilename = "/tmp/logging-sample.event.log"

var eventLog *log.Logger

func init() {
	f, err := os.OpenFile(EventLogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	eventLog = log.New(f, "", 0)
}

type EventData struct {
	Category string    `json:"category"`
	Action   string    `json:"action"`
	Label    string    `json:"label"`
	Value    string    `json:"value"`
	Time     time.Time `json:"time"`
}

func (d EventData) String() string {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(d); err != nil {
		return err.Error()
	}
	return buf.String()
}

func EventLog(d EventData) {
	eventLog.Print(d.String())
}

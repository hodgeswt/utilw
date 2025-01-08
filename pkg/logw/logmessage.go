package logw

import (
	"encoding/json"
	"fmt"
	"time"
)

type LogMessage struct {
	ProgramName string `json:"programName"`
	LogLevel    string `json:"logLevel"`
	Message     string `json:"message"`
	Time        string `json:"time"`
}

func NewLogMessage(programName string, logLevel string, message string, time string) *LogMessage {
	return &LogMessage{
		ProgramName: programName,
		LogLevel:    logLevel,
		Message:     message,
		Time:        time,
	}
}

func (it *LogMessage) json() string {
	b, err := json.Marshal(it)

	if err != nil {
		return fmt.Sprintf(`{"logLevel": "ERROR", "message": "%s", "time": "%s"}`, err.Error(), time.Now().Format(time.RFC3339))
	}

	return string(b)
}

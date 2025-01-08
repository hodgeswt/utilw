package logw

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DEBUG = 1
	INFO  = 2
	WARN  = 4
	ERROR = 8
)

const i_DEFAULT_LEVEL = ERROR
const str_DEFAULT_LEVEL = "ERROR"

type Logger struct {
	logLevel    uint
	filePath    string
	jsonLogging bool
	parsed      bool
	programName string
}

func NewLogger(programName string, logConfig *LogConfig) (*Logger, error) {
	var l = new(Logger)

	e := l.LoadConfig(logConfig)

	if l.programName == "" {
		l.programName = programName
	}

	return l, e
}

func (it *Logger) LoadConfig(logConfig *LogConfig) error {
	if it.parsed {
		return nil
	}

	// If we're not given a config,
	// attempt to load it from the
	// local environment
	if logConfig == nil {
		logConfig = loadFromEnvironment()
	}

	l := parseLogLevel(logConfig.LogLevel)
	it.logLevel = l

	it.filePath = logConfig.FilePath
	if len(it.filePath) != 0 {
		setOutFile(it.filePath)
	}

	it.jsonLogging = logConfig.JsonLogging

	return nil
}

func parseLogLevel(logLevel string) uint {
	var out uint
	switch strings.ToLower(logLevel) {
	case "all":
		out = DEBUG | INFO | WARN | ERROR
		break
	case "debug":
		out = DEBUG | INFO | WARN | ERROR
		break
	case "info":
		out = INFO | WARN | ERROR
		break
	case "warn":
		out = WARN | ERROR
		break
	case "error":
	default:
		out = ERROR
	}
	return out
}

func setOutFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)
	return logFile, nil
}

func (it *Logger) buildMessage(logLevel string, message string, v ...any) string {
	formatted := fmt.Sprintf(message, v...)
	var outMessage string

	if it.jsonLogging {
		outMessage = NewLogMessage(it.programName, logLevel, formatted, time.Now().Format(time.RFC3339)).json()
	} else {
		outMessage = fmt.Sprintf("~"+it.programName+"~["+logLevel+"] - "+message, v...)
	}

	return outMessage
}

func (it *Logger) Debugf(message string, v ...any) {
	if it.logLevel&1 != 1 {
		return
	}

	outMessage := it.buildMessage("DEBUG", message, v...)
	log.Print(outMessage)
}

func (it *Logger) Infof(message string, v ...any) {
	if (it.logLevel>>1)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("INFO", message, v...)
	log.Print(outMessage)
}

func (it *Logger) Warnf(message string, v ...any) {
	if (it.logLevel>>2)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("WARN", message, v...)
	log.Print(outMessage)
}

func (it *Logger) Errorf(message string, v ...any) {
	if (it.logLevel>>3)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("ERROR", message, v...)
	log.Print(outMessage)
}

func (it *Logger) Debug(message string) {
	if it.logLevel&1 != 1 {
		return
	}

	outMessage := it.buildMessage("DEBUG", message)
	log.Print(outMessage)
}

func (it *Logger) Info(message string) {
	if (it.logLevel>>1)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("INFO", message)
	log.Print(outMessage)
}

func (it *Logger) Warn(message string) {
	if (it.logLevel>>2)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("WARN", message)
	log.Print(outMessage)
}

func (it *Logger) Error(message string) {
	if (it.logLevel>>3)&1 != 1 {
		return
	}

	outMessage := it.buildMessage("ERROR", message)
	log.Print(outMessage)
}

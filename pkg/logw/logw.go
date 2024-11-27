package logw

import (
	"log"
	"os"
	"strings"
)

const (
	DEBUG = 1
	INFO  = 2
	WARN  = 4
	ERROR = 8
)

var level uint = ERROR
var determined = false

var loglevelw = "LOGLEVELW"

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

func determineLevel() {
	if determined {
		return
	}
	determined = true

	v, ok := os.LookupEnv(loglevelw)

	if !ok {
		level = ERROR
        return
	}

    level = parseLogLevel(v)
}

func SetLogLevel(logLevel string) {
    level = parseLogLevel(logLevel)
}

func SetOutFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    log.SetOutput(logFile)
    return logFile, nil
}

func Debugf(message string, v ...any) {
	if !determined {
		determineLevel()
	}

	if level&1 != 1 {
		return
	}

	log.Printf("[DEBUG] - "+message, v...)
}

func Infof(message string, v ...any) {
	if !determined {
		determineLevel()
	}

	if (level>>1)&1 != 1 {
		return
	}

	log.Printf("[INFO] - "+message, v...)
}

func Warnf(message string, v ...any) {
	if !determined {
		determineLevel()
	}

	if (level>>2)&1 != 1 {
		return
	}

	log.Printf("[WARN] - "+message, v...)
}

func Errorf(message string, v ...any) {
	if !determined {
		determineLevel()
	}

	if (level>>3)&1 != 1 {
		return
	}

	log.Printf("[ERROR] - "+message, v...)
}

func Debug(message string) {
	if !determined {
		determineLevel()
	}

	if level&1 != 1 {
		return
	}

	log.Print("[DEBUG] - " + message)
}

func Info(message string) {
	if !determined {
		determineLevel()
	}

	if (level>>1)&1 != 1 {
		return
	}

	log.Printf("[INFO] - " + message)
}

func Warn(message string) {
	if !determined {
		determineLevel()
	}

	if (level>>2)&1 != 1 {
		return
	}

	log.Printf("[WARN] - " + message)
}

func Error(message string) {
	if !determined {
		determineLevel()
	}

	if (level>>3)&1 != 1 {
		return
	}

	log.Printf("[ERROR] - " + message)
}

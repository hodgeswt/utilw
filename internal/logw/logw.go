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
var parsed = false

var loglevelw = "LOGLEVELW"

func parse() {
	if parsed {
		return
	}
	parsed = true

	v, ok := os.LookupEnv(loglevelw)

	if !ok {
		level = ERROR
	}

	switch strings.ToLower(v) {
	case "all":
		level = DEBUG | INFO | WARN | ERROR
		break
	case "debug":
		level = DEBUG | INFO | WARN | ERROR
		break
	case "info":
		level = INFO | WARN | ERROR
		break
	case "warn":
		level = WARN | ERROR
		break
	case "error":
	default:
		level = ERROR
	}
}

func Debugf(message string, v ...any) {
	if !parsed {
		parse()
	}

	if level&DEBUG != 1 {
		return
	}

	log.Printf("[DEBUG] - "+message, v...)
}

func Infof(message string, v ...any) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[INFO] - "+message, v...)
}

func Warnf(message string, v ...any) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[WARN] - "+message, v...)
}

func Errorf(message string, v ...any) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[ERROR] - "+message, v...)
}

func Debug(message string) {
	if !parsed {
		parse()
	}

	if level&DEBUG != 1 {
		return
	}

	log.Print("[DEBUG] - "+message)
}

func Info(message string) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[INFO] - "+message)
}

func Warn(message string) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[WARN] - "+message)
}

func Error(message string) {
	if !parsed {
		parse()
	}

	if level&INFO != 1 {
		return
	}

	log.Printf("[ERROR] - "+message)
}

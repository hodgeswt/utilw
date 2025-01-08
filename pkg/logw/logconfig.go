package logw

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

type LogConfig struct {
	LogLevel    string `json:"logLevel"`
	FilePath    string `json:"filePath"`
	JsonLogging bool   `json:"jsonLogging"`
	ProgramName string `json:"programName"`
}

func NewLogConfig() *LogConfig {
	return &LogConfig{}
}

func loadFromEnvironmentVariables() *LogConfig {
	logLevel, ok := os.LookupEnv("LOGLEVELW")

	if !ok {
		logLevel = str_DEFAULT_LEVEL
	}

	filePath, _ := os.LookupEnv("LOGWFILEPATH")

	strJsonLogging, _ := os.LookupEnv("LOGWJSONLOG")
	jsonLogging, _ := strconv.ParseBool(strJsonLogging)

	programName, _ := os.LookupEnv("LOGWPROGRAM")

	return &LogConfig{
		LogLevel:    logLevel,
		FilePath:    filePath,
		JsonLogging: jsonLogging,
		ProgramName: programName,
	}
}

func loadFromEnvironment() *LogConfig {
	filePath := "logw.json"

	globalFilePath, ok := os.LookupEnv("LOGWCONFIG")
	if ok {
		filePath = globalFilePath
	}

	j, err := os.Open(filePath)

	if err != nil {
		return loadFromEnvironmentVariables()
	}

	defer j.Close()

	b, err := io.ReadAll(j)

	if err != nil {
		return loadFromEnvironmentVariables()
	}

	var logConfig LogConfig
	err = json.Unmarshal(b, &logConfig)

	if err != nil {
		return nil
	}

	return &logConfig
}

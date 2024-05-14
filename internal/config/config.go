package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const (
	debug  = "debug"
	info   = "info"
	warn   = "warn"
	error  = "error"
	dPanic = "dpanic"
	panic  = "panic"
	fatal  = "fatal"
)

type config struct {
	GrpcPort          string        `json:"gRPC_port"`
	HttpPort          string        `json:"HTTP_port"`
	LogFile           string        `json:"path_to_log_file"`
	LogLevel          string        `json:"log_level"`
	StrBuilderBaseCap int           `json:"string_builder_base_cap"`
	UpdatingFreq      time.Duration `json:"db_updating_freq"`
}

var globalConfig config

func Read(pathToConfig string) {

	configFile, err := os.ReadFile(pathToConfig)
	if err != nil {
		log.Fatalf("failed to open config file: %s", err)
	}

	err = json.Unmarshal(configFile, &globalConfig)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %s", err)
	}

	switch globalConfig.LogLevel {
	case debug, info, warn, error, dPanic, panic, fatal:
	default:
		log.Fatalf("unknown log level in config file: %s", globalConfig.LogLevel)
	}
}

func GetGrpcPort() string {
	return globalConfig.GrpcPort
}
func GetHttpPort() string {
	return globalConfig.HttpPort
}
func GetLogFile() string {
	return globalConfig.LogFile
}
func GetLogLevel() string {
	return globalConfig.LogLevel
}
func GetStrBuilderBaseCap() int {
	return globalConfig.StrBuilderBaseCap
}
func GetUpdatingFreq() time.Duration {
	return globalConfig.UpdatingFreq * time.Second
}

package config

import (
	"io"
	"os"
)

type Config struct {
	Address       string
	User          string
	Password      string
	Debug         bool
	LogPrefix     string
	CaptureStatus bool
	StatusWriter  *io.Writer
}

func NewFromEnv(prefix string) *Config {
	_, isDebugEnabled := os.LookupEnv(prefix + "DEBUG")
	return &Config{
		Address:       os.Getenv(prefix + "ADDRESS"),
		User:          os.Getenv(prefix + "USER"),
		Password:      os.Getenv(prefix + "PASSWORD"),
		Debug:         isDebugEnabled,
		LogPrefix:     os.Getenv(prefix + "LOGPREFIX"),
		CaptureStatus: false,
		StatusWriter:  nil,
	}
}

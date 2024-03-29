package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/mymmrac/mymm-web/server/logger"
)

func LoadConfig(filename string) (Config, error) {
	var config Config

	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}

type Config struct {
	Log            Log
	Infrastructure Infrastructure
	Settings       Settings
}

type Log struct {
	Level       string
	Destination string
	Filename    string
}

type Infrastructure struct {
	HTTPServer HTTPServer
	MongoDB    MongoDB
}

type HTTPServer struct {
	Port         string
	Timeout      TextDuration
	CORSAllowAll bool
}

type MongoDB struct {
	Host    string
	Timeout TextDuration
}

type Settings struct {
	UsersFilename       string
	AuthCacheExpiration TextDuration
	CPUReadDuration     TextDuration
}

const (
	logDestinationStdout = "stdout"
	logDestinationStderr = "stderr"
	logDestinationFile   = "file"
)

const (
	logLevelError = "error"
	logLevelWarn  = "warn"
	logLevelInfo  = "info"
	LogLevelDebug = "debug"
)

func (c Config) ConfigureLogger(log *logger.Log) error {
	switch c.Log.Destination {
	case logDestinationStdout:
		log.SetOutput(os.Stdout)
	case logDestinationStderr:
		log.SetOutput(os.Stderr)
	case logDestinationFile:
		if err := log.SetOutputFile(c.Log.Filename); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unkown logger destination: %q", c.Log.Destination)
	}

	switch c.Log.Level {
	case logLevelError, logLevelWarn, logLevelInfo, LogLevelDebug:
		log.SetLevel(c.Log.Level)
	default:
		return fmt.Errorf("unkown logger level: %q", c.Log.Level)
	}

	return nil
}

type TextDuration struct {
	time.Duration
}

func (d *TextDuration) UnmarshalText(text []byte) error {
	duration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}

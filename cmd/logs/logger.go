package logs

import (
	"errors"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func NewLogger() {
	var err error
	var logFilePath = "/tmp/log/audiofile.json"
	cfg := zap.Config{
		Level:       zap.NewAtomicLevel(),
		Encoding:    "json",
		OutputPaths: []string{logFilePath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "file",
			StacktraceKey:  "stacktrace",
			EncodeName:     zapcore.FullNameEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
		},
	}
	if err := createLogFileIfNotExists(logFilePath); err != nil {
		panic(err)
	}
	Logger, err = cfg.Build()
	if err != nil {
		println("error: setup the logger")
		panic(err)
	}
	defer Logger.Sync()
	Logger.Info("Logger building succeeded")
}

func createLogFileIfNotExists(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		dirName := filepath.Dir(path)
		if _, err = os.Stat(dirName); err != nil {
			if err = os.MkdirAll(dirName, os.FileMode(0777)); err != nil {
				panic(err)
			}
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

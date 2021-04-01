package utils

import (
	"testing"
	"time"
)

func setup() {
	c := LoggerConfig{
		Filename:   "../test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
	}
	InitLogger(&c)
}

func TestLog(t *testing.T) {
	// XLogger("jejeedd")
	setup()
	Logger.Info("test for ")
}

// XLogger return text
/*
test xlogger
*/
func TestXLogger(t *testing.T) {
	text := "test2"
	setup()
	// defer Logger.Sync() // flushes buffer, if any
	Logger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", text,
		"attempt", 3,
		"backoff", time.Second,
	)
	Logger.Infof("Failed to fetch URL: %s", text)
}

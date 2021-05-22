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
func teardown() {

}

func TestLog(t *testing.T) {
	// XLogger("jejeedd")
	Logger.Info("test for ")
}

// XLogger return text
/*
test xlogger
*/
func TestXLogger(t *testing.T) {
	text := "test2"
	// defer Logger.Sync() // flushes buffer, if any
	Logger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", text,
		"attempt", 3,
		"backoff", time.Second,
	)
	Logger.Infof("Failed to fetch URL: %s", text)
}

func TestFoo(t *testing.T) {
	setup()
	t.Run("TestXLogger", TestXLogger)
	t.Run("TestXLogger", TestLog)
	teardown()
}

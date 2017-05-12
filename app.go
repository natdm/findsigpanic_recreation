package main

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	jsonLogger := log.NewJSONLogger(os.Stdout)
	jsonLogger = log.WithPrefix(jsonLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller, "service", "api")
	jsonLogger.Log("message", "started")
}

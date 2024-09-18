package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	logger.Info("User logged in",
		zap.String("username", "LUMOSHIVE-ACADEMY"),
		zap.Int("user_id", 12345),
	)

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.DPanic("This is a DPanic message")
	logger.Panic("This is a Panic message")
	logger.Fatal("This is a Fatal message")
}

// package main

// import (
// 	"os"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// func main() {
// 	file, _ := os.Create("app.log")
// 	writer := zapcore.AddSync(file)
// 	core := zapcore.NewCore(
// 		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
// 		writer,
// 		zap.InfoLevel,
// 	)
// 	logger := zap.New(core)
// 	defer logger.Sync()

// 	logger.Info("This is a log message to a file")
// }

package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Menggunakan JSON Encoder
	jsonEncoderConfig := zap.NewProductionEncoderConfig()
	jsonLogger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(jsonEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	))
	defer jsonLogger.Sync()
	jsonLogger.Info("This is a JSON formatted log")

	// Menggunakan Console Encoder
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleLogger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	))
	defer consoleLogger.Sync()
	consoleLogger.Info("This is a console formatted log")
}

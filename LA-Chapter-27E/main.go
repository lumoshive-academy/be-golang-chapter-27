// package main

// import (
// 	"os"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// func main() {
// 	// Konfigurasi encoder untuk output JSON
// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.TimeKey = "timestamp"
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

// 	// Sink untuk stdout
// 	writer := zapcore.AddSync(os.Stdout)

// 	// Core dengan filter level log minimal
// 	core := zapcore.NewCore(
// 		zapcore.NewJSONEncoder(encoderConfig),
// 		writer,
// 		zap.WarnLevel, // Hanya log dengan level Warn dan di atasnya yang akan dicatat
// 	)

// 	// Membuat logger
// 	logger := zap.New(core)
// 	defer logger.Sync()

// 	logger.Debug("This is a debug message")  // Tidak akan dicatat
// 	logger.Info("This is an info message")   // Tidak akan dicatat
// 	logger.Warn("This is a warning message") // Akan dicatat
// 	logger.Error("This is an error message") // Akan dicatat
// }

package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Menambahkan hook yang akan dipanggil untuk setiap log entry
	hookedLogger := logger.WithOptions(zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level >= zapcore.ErrorLevel {
			// Tindakan tambahan jika level log adalah Error atau lebih tinggi
			fmt.Printf("An error log was generated: %v\n", entry.Message)
		}
		return nil
	}))

	// Logging contoh
	hookedLogger.Info("This is an informational message")
	hookedLogger.Error("This is an error message")
}

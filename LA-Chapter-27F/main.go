package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Structured Logging: Menggunakan zap untuk membuat log yang terstruktur
func main() {
	// 1. Jangan Tulis Log Sendiri: Gunakan library zap
	logger := NewLogger()
	defer logger.Sync()

	// 2. Log pada Tingkat yang Tepat: Menggunakan level log yang sesuai
	logger.Debug("Starting application", zap.String("environment", "development"))

	// 3. Hormati Kategori Log: Kelompokkan log berdasarkan kategori
	logger.Info("HTTP request received", zap.String("method", "GET"), zap.String("endpoint", "/api/v1/resource"))

	// 4. Tulis Log yang Bermakna: Sertakan informasi penting
	userID := "user123"
	logger.Warn("User login attempt failed", zap.String("userID", userID), zap.String("reason", "invalid password"))

	// 5. Log Harus dalam Bahasa Inggris: Standar industri
	logger.Error("Database connection failed", zap.String("database", "user_db"), zap.Error(err))

	// 6. Jangan Log Terlalu Banyak atau Terlalu Sedikit: Gunakan log level yang tepat
	logger.Debug("Detailed debug information", zap.Any("debug_info", someDebugData))

	// 7. Log Harus Konsisten: Gunakan format yang konsisten
	logger.Info("Transaction completed", zap.String("transactionID", "txn_456"), zap.Float64("amount", 99.99))

	// 8. Log Harus Mudah Dicari: Gunakan ID unik atau label
	logger.Info("User session started", zap.String("sessionID", "session_789"))

	// 9. Log Harus Aman: Hindari informasi sensitif
	// logger.Info("Password reset", zap.String("new_password", "supersecretpassword")) // Jangan lakukan ini
	logger.Info("Password reset", zap.String("userID", userID))

	// 10. Log Harus Mudah Dianalisis: Menggunakan alat analisis log
	// Log sudah dalam format JSON yang mudah dianalisis dengan alat seperti Elasticsearch
}

// NewLogger membuat logger baru dengan konfigurasi produksi
func NewLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel, // Set default log level to Info
	)

	logger := zap.New(core)
	return logger
}

var err = &os.PathError{Op: "open", Path: "non_existent_file", Err: os.ErrNotExist}
var someDebugData = map[string]interface{}{
	"key1": "value1",
	"key2": 42,
	"key3": true,
}

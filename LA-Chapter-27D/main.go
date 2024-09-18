// package main

// import (
// 	"go.uber.org/zap"
// )

// func main() {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()

// 	// Menambahkan field pada log
// 	logger.Info("User logged in",
// 		zap.String("username", "john_doe"),
// 		zap.Int("user_id", 12345),
// 		zap.String("role", "admin"),
// 	)

// 	// Menambahkan field dengan nilai numerik dan boolean
// 	logger.Info("Transaction completed",
// 		zap.String("transaction_id", "tx123456789"),
// 		zap.Float64("amount", 100.75),
// 		zap.Bool("successful", true),
// 	)
// }

package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Error("An error occurred", zap.String("error_detail", "unable to connect to database"))

	logger.Info("An informative log with stack trace",
		zap.String("info", "some important context"),
		zap.Stack("stacktrace"),
	)
}

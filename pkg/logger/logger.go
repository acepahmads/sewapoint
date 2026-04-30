package logger

import (
	"log"
)

// Cara pakai
// logger.Info("user login: %d", userID)
// logger.Error("failed booking: %v", err)

func Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Printf("[WARN] "+msg, args...)
}

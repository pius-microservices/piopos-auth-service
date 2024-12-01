package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() (*log.Logger, *log.Logger) {
	currentDate := time.Now().Format("2006-01-02")

	if err := os.MkdirAll("logs/error", os.ModePerm); err != nil {
		log.Fatalf("Failed to create error log directory: %v", err)
	}
	if err := os.MkdirAll("logs/debug", os.ModePerm); err != nil {
		log.Fatalf("Failed to create debug log directory: %v", err)
	}

	errorLogFilePath := filepath.Join("logs/error", fmt.Sprintf("error-%s.log", currentDate))
	debugLogFilePath := filepath.Join("logs/debug", fmt.Sprintf("debug-%s.log", currentDate))

	errorLogFile, err := os.OpenFile(errorLogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	debugLogFile, err := os.OpenFile(debugLogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Failed to open debug log file: %v", err)
	}

	errorLogger := log.New(errorLogFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
	debugLogger := log.New(debugLogFile, "DEBUG: ", log.LstdFlags|log.Lshortfile)

	return errorLogger, debugLogger
}

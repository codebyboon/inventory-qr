package logger

import (
	"log"
	"os"
	"path/filepath"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {

	// Paths for the log files
	infoLogPath := "./logs/info.log"
	errorLogPath := "./logs/error.log"

	infoLogDir := filepath.Dir(infoLogPath)

	if _, err := os.Stat(infoLogDir); os.IsNotExist(err) {
		if err := os.MkdirAll(infoLogDir, os.ModePerm); err != nil {
			log.Printf("Error creating directory: %v", err)
		}
	}

	errorLogDir := filepath.Dir(errorLogPath)
	if _, err := os.Stat(errorLogDir); os.IsNotExist(err) {
		if err := os.MkdirAll(errorLogDir, os.ModePerm); err != nil {
			log.Printf("Error creating directory: %v", err)
		}
	}

	// Create/Open log files
	infoLogFile, err := os.OpenFile(infoLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening/creating info log file: %v", err)
	}
	errorLogFile, err := os.OpenFile(errorLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening/creating error log file: %v", err)
	}

	// Initialize the InfoLogger to write to infoLogFile
	InfoLogger = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Initialize the ErrorLogger to write to errorLogFile
	ErrorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

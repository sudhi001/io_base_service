package io_base_service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Logger is a custom logger without logrus dependency
type Logger struct {
	logFile *os.File
	fields  map[string]interface{}
}

// Global logger instance
var AppLogger *Logger

// InitLogger initializes the logger with file-based logging
func InitLogger(logFolder string, logLevel string) {
	// Ensure log folder exists
	if err := os.MkdirAll(logFolder, os.ModePerm); err != nil {
		fmt.Println("Error creating log directory:", err)
	}

	// Create log file with date-wise naming
	logFileName := logFolder + "/" + time.Now().Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file, logging to stdout:", err)
		AppLogger = &Logger{logFile: os.Stdout}
	} else {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
		AppLogger = &Logger{logFile: logFile}
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// log writes structured log messages
func (l *Logger) log(level string, message string, fields map[string]interface{}) {
	if l == nil {
		fmt.Println("Logger not initialized")
		return
	}
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["level"] = level
	fields["time"] = time.Now().Format(time.RFC3339)
	fields["message"] = message

	logData, err := json.Marshal(fields)
	if err != nil {
		fmt.Println("Error marshaling log data:", err)
		return
	}

	log.Println(string(logData))
}

// Logging helper methods
func Info(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	AppLogger.log("info", msg, nil)
}

func Warn(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	AppLogger.log("warning", msg, nil)
}

func Error(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	AppLogger.log("error", msg, nil)
}

func Debug(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	AppLogger.log("debug", msg, nil)
}

func Fatalf(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	AppLogger.log("fatal", msg, nil)
	fmt.Println("FATAL:", msg) // Print to stdout in case logging fails
	os.Exit(1)                 // Exit the program with a failure code
}

// WithFields adds fields to the log message
func WithFields(fields map[string]interface{}) *Logger {
	return &Logger{fields: fields}
}

// WithError logs an error with a message
func WithError(err error) *Logger {
	return &Logger{fields: map[string]interface{}{"error": err.Error()}}
}

// Log with fields
func (l *Logger) Info(message string) {
	AppLogger.log("info", message, l.fields)
}

func (l *Logger) Warn(message string) {
	AppLogger.log("warning", message, l.fields)
}

func (l *Logger) Error(message string) {
	AppLogger.log("error", message, l.fields)
}

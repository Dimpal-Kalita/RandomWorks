package main

import (
	"os"

	logger "github.com/Dimpal-Kalita/RandomWorks/logger/utils"
)

func main() {
	consoleLogger := logger.NewLogger(os.Stdout, logger.INFO)
	consoleLogger.Info("info message")
	consoleLogger.Warn("warn message")
	consoleLogger.Error("error message")
	consoleLogger.Debug("debug message")

	// Crating a log file
	file, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		consoleLogger.Error("Failed to open log file")
		os.Exit(1)
	}
	defer file.Close()
	fileLogger := logger.NewLogger(file, logger.DEBUG)
	fileLogger.Info("info message [Not visible at debug level]")
	fileLogger.Warn("warn message [Not visible at debug level]")
	fileLogger.Error("error message [Not visible at debug level]")
	fileLogger.Debug("debug message [Visible at debug level]")

}

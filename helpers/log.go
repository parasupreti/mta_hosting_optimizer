package helpers

import (
	"log"
	"os"
	"time"
)

func printLog(debugLevel string, args ...interface{}) {
	currentTime := time.Now()
	path := "tmp/logs/logfile-"
	fileName := path + currentTime.Format("2006-01-02") + ".log"

	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger := log.New(logFile, debugLevel, log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println(args...)

	defer logFile.Close()
}

// InfoLog info log errors
func InfoLog(args ...interface{}) {
	printLog("INFO: ", args...)
}

// ErrorLog function to log errors.
func ErrorLog(args ...interface{}) {
	printLog("ERROR: ", args...)
}

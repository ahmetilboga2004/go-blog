package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel string

const (
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

var (
	loggers map[LogLevel]*log.Logger
	logFile *os.File
	today   string
)

func createLogFile() error {
	if logFile != nil {
		logFile.Close()
	}

	today = time.Now().Format("2006-01-02")
	var err error
	logFile, err = os.OpenFile("logs_"+today+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	loggers = make(map[LogLevel]*log.Logger)
	loggers[INFO] = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	loggers[WARNING] = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime)
	loggers[ERROR] = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)

	return nil
}

func Log(level LogLevel, format string, args ...interface{}) {
	currentDate := time.Now().Format("2006-01-02")
	if currentDate != today {
		if err := createLogFile(); err != nil {
			log.Printf("Yeni log dosyası oluşturulamadı: %v", err)
			return
		}
	}

	message := fmt.Sprintf(format, args...)

	if logger, exists := loggers[level]; exists {
		logger.Println(message)
	}
}

func init() {
	if err := createLogFile(); err != nil {
		log.Fatal("Log dosyası açılamadı: ", err)
	}
}

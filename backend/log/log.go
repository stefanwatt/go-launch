package log

import (
	Config "go-launch/backend/config"
	"log"
	"os"
	"path"
	"reflect"
)

var Print func(v ...interface{})

func InitLogger() {
	// delete app.log
	logfile := path.Join(Config.HOME, "go-launch.log")
	err := os.Remove(logfile)
	if err != nil {
		log.Println("Log file not found, creating new one")
	}
	logFile, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(logFile)
	// setup timestapms for log entries
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	Print = log.Println
}

func PrintArray[T any](arr []T) {
	for _, value := range arr {
		Print(value)
	}
}

func PrintProperties[T any](arr []T, fieldName string) {
	for _, item := range arr {
		value := reflect.ValueOf(item).FieldByName(fieldName)
		Print(value)
	}
}

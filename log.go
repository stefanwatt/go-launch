package main

import (
	"log"
	"os"
	"reflect"
)

var print func(v ...interface{})

func initLogger() {
	// delete app.log
	err := os.Remove("app.log")
	if err != nil {
		log.Println("Log file not found, creating new one")
	}
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(logFile)
	// setup timestapms for log entries
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	print = log.Println
}

func printArray[T any](arr []T) {
	for _, value := range arr {
		print(value)
	}
}

func printProperties[T any](arr []T, fieldName string) {
	for _, item := range arr {
		value := reflect.ValueOf(item).FieldByName(fieldName)
		print(value)
	}
}

package util

import (
	"log"
	"os"
)

var logFile *os.File

func InitLogger() {
	file, err := os.OpenFile("./log/data.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	logFile = file
}

func DisposeLogger() {
	logFile.Close()
}

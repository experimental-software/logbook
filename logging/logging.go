package logging

import (
	"log"
	"os"
)

var (
	warningLog *log.Logger
	infoLog    *log.Logger
	errorLog   *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLog = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	infoLog.Println(message)
}

func Warn(message string) {
	warningLog.Println(message)
}

func Error(message string, err error) {
	errorLog.Println("ERROR: " + message)
	errorLog.Println(err.Error())
}

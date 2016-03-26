package logger

// Support Log Level for logging
// Log Level can be changed runtime
import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	CRITICAL
	FATAL
)

var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

// Initialize logger
// if filename is empty then logs are written to standard output
// default Log Level is DEBUG
func Init(fileName string) error {
	var err error
	fileHandle := os.Stderr
	logLevel = DEBUG

	if fileName != "" {
		log.Printf("Writting logs to file:%s\n", fileName)
		fileHandle, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if nil != err {
			log.Print(err)
			fileHandle = os.Stderr
		}
	}
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)

	// also set default log output to this file handle
	log.SetOutput(fileHandle)
	return nil
}

// Set Log Level
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// log the input logging msg into initialized file handle
func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

// debug logging
func Debug(format string, v ...interface{}) {
	if DEBUG >= logLevel {
		llog("[DEBUG] "+format, v...)
	}
}

// info logging
func Info(format string, v ...interface{}) {
	if INFO >= logLevel {
		llog("[INFO] "+format, v...)
	}
}

// warn logging
func Warn(format string, v ...interface{}) {
	if WARN >= logLevel {
		llog("[WARN] "+format, v...)
	}
}

// error logging
func Error(format string, v ...interface{}) {
	if ERROR >= logLevel {
		llog("[ERROR] "+format, v...)
	}
}

// fatal logging
func Fatal(format string, v ...interface{}) {
	if FATAL >= logLevel {
		llog("[FATAL] "+format, v...)
	}
}

// critical logging
// NOTE: process terminate on critical logging
func Critical(format string, v ...interface{}) {
	if CRITICAL >= logLevel {
		llog("[CRITICAL] "+format, v...)
		os.Exit(1)
	}
}

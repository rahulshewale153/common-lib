package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logLevel int

const (
	TRACE logLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	CRITICAL
	FATAL
)

type Logger struct {
	logger    *log.Logger
	logLevel  logLevel
	calldepth int
}

// New creates a new logger with default settings.
// By default, logs are written to standard error.
func NewLogger() *Logger {
	l := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds)
	logger := Logger{logger: l,
		logLevel:  ERROR,
		calldepth: 3}
	return &logger
}

// Output sets the output destination of the logger.
// By default, logs are written to standard error.
func (logger *Logger) SetOutput(out io.Writer) {
	logger.logger.SetOutput(out)
}

func (logger *Logger) SetLogLevel(level logLevel) {
	logger.logLevel = level
}

func (logger *Logger) setCalldepth(depth int) {
	logger.calldepth = depth
}

func (logger *Logger) SetOutputFile(fileName string) {
	fileHandle, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Print(err)
		fileHandle = os.Stderr
	}
	logger.SetOutput(fileHandle)
}

func (logger *Logger) llog(format string, v ...interface{}) {
	logger.logger.Output(logger.calldepth, fmt.Sprintf(format, v...))
}

func (logger *Logger) Trace(format string, v ...interface{}) {
	if TRACE >= logger.logLevel {
		logger.llog("[TRACE] "+format, v...)
	}
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	if DEBUG >= logger.logLevel {
		logger.llog("[DEBUG] "+format, v...)
	}
}

func (logger *Logger) Info(format string, v ...interface{}) {
	if INFO >= logger.logLevel {
		logger.llog("[INFO] "+format, v...)
	}
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	if WARN >= logger.logLevel {
		logger.llog("[WARN] "+format, v...)
	}
}

func (logger *Logger) Error(format string, v ...interface{}) {
	if ERROR >= logger.logLevel {
		logger.llog("[ERROR] "+format, v...)
	}
}

func (logger *Logger) Critical(format string, v ...interface{}) {
	if CRITICAL >= logger.logLevel {
		logger.llog("[CRITICAL] "+format, v...)
	}
}

func (logger *Logger) Fatal(format string, v ...interface{}) {
	if FATAL >= logger.logLevel {
		logger.llog("[FATAL] "+format, v...)
		os.Exit(1)
	}
}

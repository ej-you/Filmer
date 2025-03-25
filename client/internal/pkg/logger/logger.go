package logger

import (
	"log"
	"os"
	"sync"
)

// Logger interface
type Logger interface {
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
}

// Logger implementation
type appLogger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	errorLog *log.Logger
}

var once sync.Once
var logger = new(appLogger)

// Logger constructor
func NewLogger() Logger {
	once.Do(func() {
		logger.debugLog = log.New(os.Stdout, "[DEBUG]\t", log.Ldate|log.Ltime)
		logger.infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
		logger.errorLog = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)
	})
	return logger
}

func (l appLogger) Debug(args ...any) {
	l.debugLog.Print(args...)
}

func (l appLogger) Debugf(template string, args ...any) {
	l.debugLog.Printf(template, args...)
}

func (l appLogger) Info(args ...any) {
	l.infoLog.Println(args...)
}

func (l appLogger) Infof(template string, args ...any) {
	l.infoLog.Printf(template, args...)
}

func (l appLogger) Error(args ...any) {
	l.errorLog.Println(args...)
}

func (l appLogger) Errorf(template string, args ...any) {
	l.errorLog.Printf(template, args...)
}

func (l appLogger) Fatal(args ...any) {
	l.errorLog.Fatal(args...)
}

func (l appLogger) Fatalf(template string, args ...any) {
	l.errorLog.Fatalf(template, args...)
}

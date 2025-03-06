package logger

import (
	"log"
	"os"
	"sync"
)


// интерфейс логера
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


// структура логера
type appLogger struct {
	debugLog *log.Logger
	infoLog *log.Logger
	errorLog *log.Logger
}


var once sync.Once
var logger = new(appLogger)

// конструктор для типа интерфейса Logger
func NewLogger() Logger {
	once.Do(func() {
		logger.debugLog = log.New(os.Stdout, "[DEBUG]\t", log.Ldate|log.Ltime)
		logger.infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
		logger.errorLog = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)
	})
	return logger
}

func (this appLogger) Debug(args ...any) {
	this.debugLog.Print(args...)
}

func (this appLogger) Debugf(template string, args ...any) {
	this.debugLog.Printf(template, args...)
}

func (this appLogger) Info(args ...any) {
	this.infoLog.Println(args...)
}

func (this appLogger) Infof(template string, args ...any) {
	this.infoLog.Printf(template, args...)
}

func (this appLogger) Error(args ...any) {
	this.errorLog.Println(args...)
}

func (this appLogger) Errorf(template string, args ...any) {
	this.errorLog.Printf(template, args...)
}

func (this appLogger) Fatal(args ...any) {
	this.errorLog.Fatal(args...)
}

func (this appLogger) Fatalf(template string, args ...any) {
	this.errorLog.Fatalf(template, args...)
}

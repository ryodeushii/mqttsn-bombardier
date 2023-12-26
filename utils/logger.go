package utils

import "fmt"

type ILogger interface {
	Info(message string, data ...interface{})
	Warn(message string, data ...interface{})
	Error(message string, data ...interface{})
	Panic(message string, data ...interface{})
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(message string, data ...interface{}) {
	println(fmt.Sprintf("[log] %s  %+v", message, data))
}

func (l *Logger) Warn(message string, data ...interface{}) {
	println(fmt.Sprintf("[log] %s  %+v", message, data))
}

func (l *Logger) Error(message string, data ...interface{}) {
	// println(fmt.Sprintf("[log] %s  %+v", message, data))
}

func (l *Logger) Panic(message string, data ...interface{}) {
	msg := fmt.Sprintf("[log] %s  %+v", message, data)
	println(msg)
	panic(msg)
}

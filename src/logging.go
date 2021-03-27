package main

import (
	"io"
	"log"
)

const (
	llNone    = 0
	llInfo    = 1
	llWarning = 2
	llError   = 3
)

type Logger struct {
	wl *log.Logger
	il *log.Logger
	el *log.Logger

	ll int
}

func NewLogger(infoHandle io.Writer, warningHandle io.Writer, errHandle io.Writer, logLevel int) *Logger {
	return &Logger{
		wl: log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime),
		il: log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime),
		el: log.New(errHandle, "ERROR: ", log.Ldate|log.Ltime),

		ll: logLevel,
	}
}

func (l *Logger) SetLevel(level int) {
	l.ll = level
}

func (l *Logger) Warning(a ...interface{}) {
	if l.ll >= llWarning {
		l.wl.Println(a...)
	}
}

func (l *Logger) Error(a ...interface{}) {
	if l.ll >= llError {
		l.el.Println(a...)
	}
}

func (l *Logger) Info(a ...interface{}) {
	if l.ll >= llInfo {
		l.il.Println(a...)
	}
}

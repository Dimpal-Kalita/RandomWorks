package logger

import (
	"io"
	"sync"
	"time"
)




type loglevel int

const (
	INFO loglevel = iota
	WARN
	ERROR
	DEBUG
)

type Logger struct {
	mu sync.Mutex
	level loglevel
	output io.Writer
}

func NewLogger(output io.Writer, level loglevel) *Logger {
	return &Logger{
		output: output,
		level: level,
	}
}

func (l *Logger) log(level loglevel, msg string) error{
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level {
		return nil
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := [...]string{"INFO", "WARN", "ERROR", "DEBUG"}[level]
	_, err := l.output.Write([]byte(timestamp + " [" + levelStr + "] " + msg + "\n"))
	return err
}

func (l *Logger) Info(msg string) error {
	return l.log(INFO, msg)
}
func (l *Logger) Warn(msg string) error {
	return l.log(WARN, msg)
}
func (l *Logger) Error(msg string) error {
	return l.log(ERROR, msg)
}
func (l *Logger) Debug(msg string) error {
	return l.log(DEBUG, msg)
}
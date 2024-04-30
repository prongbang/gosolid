package ocp

import "fmt"

type Logger interface {
	Log(msg string)
}

type FileLogger struct{}

func (l *FileLogger) Log(msg string) {
	fmt.Println(" > Log to file:", msg)
}

func NewFileLogger() Logger {
	return &FileLogger{}
}

type DatabaseLogger struct{}

func (l *DatabaseLogger) Log(msg string) {
	fmt.Println(" > Log to database:", msg)
}

func NewDatabaseLogger() Logger {
	return &DatabaseLogger{}
}

func Example() {
	fmt.Println("Open/Closed Principle")

	fLog := NewFileLogger()
	dLog := NewDatabaseLogger()
	fLog.Log("Devไปวันๆ")
	dLog.Log("Devไปวันๆ")
}

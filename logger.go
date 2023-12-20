package sdk

import "io"

type Logger interface {
	// Set the logger level
	SetLevel(level string)

	Level() string

	// Set the logger output
	SetOutput(w io.Writer)

	Output() io.Writer

	// Info logs to info level
	Info(message string, keyvals ...interface{})
	// Debug logs to debug level
	Debug(message string, keyvals ...interface{})
	// Error logs to error level
	Error(message string, keyvals ...interface{})
	// Warn logs to warn level
	Warn(message string, keyvals ...interface{})
	// Trace logs to trace level
	Trace(message string, keyvals ...interface{})

	StandardWriter() io.Writer

	IsInfo() bool
	IsDebug() bool
	IsError() bool
	IsTrace() bool
	IsWarn() bool
}

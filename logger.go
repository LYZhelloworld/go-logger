package logger

import (
	"io"
	"os"
)

type Logger interface {
	// Enabled returns if the specific level is enabled.
	Enabled(level Level) bool

	// Trace outputs log in TRACE level.
	Trace(msg string)

	// Debug outputs log in DEBUG level.
	Debug(msg string)

	// Info outputs log in INFO level.
	Info(msg string)

	// Warn outputs log in WARN level.
	Warn(msg string)

	// Error outputs log in ERROR level.
	Error(msg string)

	// Fatal outputs log in FATAL level.
	Fatal(msg string)

	// WithField attaches a key/value pair to the logger.
	WithField(key string, value interface{}) Logger

	// WithError attaches an error to the logger.
	WithError(err error) Logger

	// WithLazyField attaches a key/value pair which is lazy-loaded to the logger.
	WithLazyField(key string, value func() interface{}) Logger
}

// GetLoggerWithConfig gets logger with specified config.
func GetLoggerWithConfig(writer io.Writer, level Level) Logger {
	return &loggerImpl{
		lowestLevel: level,
		w:           writer,
		dataField:   nil,
		prev:        nil,
	}
}

var defaultLogger = GetLoggerWithConfig(os.Stderr, Trace)

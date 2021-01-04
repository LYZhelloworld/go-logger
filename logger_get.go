package logger

import "os"

// GetDefaultLogger gets root logger with default config.
func GetDefaultLogger() Logger {
	return defaultLogger
}

// GetNopLogger gets a dummy logger, toggling off all logging functions.
func GetNopLogger() Logger {
	return &nopLoggerImpl{}
}

// GetFileLogger gets a logger that will write logs to a file, returning the logger
// and a function to close the file before the application stops.
func GetFileLogger(filename string, level Level) (logger Logger, close func()) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	close = func() {
		_ = f.Close()
	}
	logger = GetLoggerWithConfig(f, level)
	return
}

package logger

// Enabled returns if the specific level is enabled.
func Enabled(level Level) bool {
	return globalLogger.Enabled(level)
}

// Trace outputs log in TRACE level.
func Trace(msg string) {
	globalLogger.Trace(msg)
}

// Debug outputs log in DEBUG level.
func Debug(msg string) {
	globalLogger.Debug(msg)
}

// Info outputs log in INFO level.
func Info(msg string) {
	globalLogger.Info(msg)
}

// Warn outputs log in WARN level.
func Warn(msg string) {
	globalLogger.Warn(msg)
}

// Error outputs log in ERROR level.
func Error(msg string) {
	globalLogger.Error(msg)
}

// Fatal outputs log in FATAL level.
func Fatal(msg string) {
	globalLogger.Fatal(msg)
}

// WithField attaches a key/value pair to the logger.
func WithField(key string, value interface{}) Logger {
	return globalLogger.WithField(key, value)
}

// WithError attaches an error to the logger.
func WithError(err error) Logger {
	return globalLogger.WithError(err)
}

// WithLazyField attaches a key/value pair which is lazy-loaded to the logger.
func WithLazyField(key string, value func() interface{}) Logger {
	return globalLogger.WithLazyField(key, value)
}

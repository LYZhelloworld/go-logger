package logger

// Level is the log level.
type Level int

const (
	// LevelTrace is the lowest log level used for tracing the flow of execution.
	LevelTrace Level = iota
	// LevelDebug is used for debugging.
	LevelDebug
	// LevelInfo is used to highlight the progress of application, information and other useful data.
	LevelInfo
	// LevelWarn is for potentially harmful message but not considered an error usually.
	LevelWarn
	// LevelError means that some error has happened, and the application may stop or still continue running.
	LevelError
	// LevelFatal is usually for those most severe errors that cause the application to abort immediately.
	LevelFatal
)

# go-logger
Go-logger is a logger implemented in Go.

## Usage
```go
var log Logger

// Get default logger:
log = logger.GetDefaultLogger()

// Get file logger with filename:
log = logger.GetFileLogger(filename, level)

// Get logger with a writer and a level:
log = logger.GetLoggerWithConfig(writer, level)

// Set global logger:
logger.SetGlobalLogger(log)

// You can then get the global logger when you need to use it:
log = logger.GetGlobalLogger()

// Or use the global functions. They will call global logger:
logger.Trace("trace")
logger.Debug("debug")
logger.Info("info")
logger.Warn("warn")
logger.Error("error")
logger.Fatal("fatal")

// Attach a field when logging:
log.WithField("data", data).Info("received data")

// Attach an error:
log.WithError(err).Error("an error occurred")

// Attach a field that needs to be lazy-loaded when logging:
log.WithLazyField("field", func() interface{} { return DoSomething() }).Info("result")
```
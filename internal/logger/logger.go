package logger

var Log Logger

// logger interface based zap's exported method
// since this is an interface, we can also swap to other log libraries if they also support log levels, and structured logging with fields
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync() error // Ensures all buffered logs are flushed.
}

// SetLogger set to exported to make it easier when we want to mock the logger during tests
func SetLogger(logger Logger) {
	Log = logger
}

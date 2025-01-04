package logging

import (
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger is an interface that defines methods for logging information, errors and warnings.
type Logger interface {
	LogInfo(message string, data ...any)
	LogError(message string, err ...any)
	LogWarn(message string, data ...any)
	LogDebug(message string, data ...any)
	SetLevel(level LoggerLevel)
	SetOutput(output *os.File)
	GetLevel() LoggerLevel
	GetLogrusEntry() *logrus.Entry
}

// LoggerLevel represents the logging level.
type LoggerLevel logrus.Level

// Used to ensure that LoggerImpl is a singleton.
var (
	once     sync.Once
	instance Logger
)

// NewLogger returns the singleton instance of LoggerImpl.
func NewLogger() Logger {
	once.Do(func() {
		instance = initLogger()
	})
	return instance
}

// initLogger initializes LoggerImpl.
func initLogger() Logger {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	consoleLogger := logrus.New()

	consoleLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	consoleLogger.SetOutput(multiWriter)

	l := &LoggerImpl{
		consoleEntry: logrus.NewEntry(consoleLogger),
	}

	return l
}

// LoggerImpl is a struct that implements the Logger interface.
type LoggerImpl struct {
	consoleEntry *logrus.Entry
}

// SetLevel sets the logging level.
func (l *LoggerImpl) SetLevel(level LoggerLevel) {
	l.consoleEntry.Logger.SetLevel(logrus.Level(level))
}

// SetOutput sets the logging output.
func (l *LoggerImpl) SetOutput(output *os.File) {
	l.consoleEntry.Logger.SetOutput(output)
}

// GetLevel returns the current logging level.
func (l *LoggerImpl) GetLevel() LoggerLevel {
	return LoggerLevel(l.consoleEntry.Logger.GetLevel())
}

// LogInfo logs an informational message.
func (l *LoggerImpl) LogInfo(message string, data ...any) {
	function, file, line := getCallerDetails()
	fields := logrus.Fields{
		"function": function,
		"file":     file,
		"line":     line,
		"data":     data,
	}

	l.consoleEntry.WithFields(fields).Info(message)
}

// LogError logs an error message along with the error if provided.
func (l *LoggerImpl) LogError(message string, err ...any) {
	function, file, line := getCallerDetails()
	fields := logrus.Fields{
		"function": function,
		"file":     file,
		"line":     line,
		"error":    err,
	}

	l.consoleEntry.WithFields(fields).Error(message)
}

// LogWarn logs a warning message along with additional data if provided.
func (l *LoggerImpl) LogWarn(message string, data ...any) {
	function, file, line := getCallerDetails()
	fields := logrus.Fields{
		"function": function,
		"file":     file,
		"line":     line,
		"data":     data,
	}

	l.consoleEntry.WithFields(fields).Warn(message)
}

// LogDebug logs a debug message along with additional data if provided.
func (l *LoggerImpl) LogDebug(message string, data ...any) {
	function, file, line := getCallerDetails()
	fields := logrus.Fields{
		"function": function,
		"file":     file,
		"line":     line,
		"data":     data,
	}

	l.consoleEntry.WithFields(fields).Debug(message)
}

// GetLogrusEntry returns the logrus entry.
func (l *LoggerImpl) GetLogrusEntry() *logrus.Entry {
	return l.consoleEntry
}

func getCallerDetails() (string, string, int) {
	pc, file, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if !ok || details == nil {
		return "unknown", "unknown", 0
	}
	return details.Name(), file, line
}

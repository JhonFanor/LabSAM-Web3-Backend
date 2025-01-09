package logging

import (
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	LogInfo(message string, data ...any)
	LogError(message string, err ...any)
	LogWarn(message string, data ...any)
	LogDebug(message string, data ...any)
	SetLevel(level LoggerLevel)
	SetOutput(output io.Writer)
	GetLevel() LoggerLevel
	GetLogrusEntry() *logrus.Entry
}

type LoggerLevel logrus.Level

const (
	LevelDebug LoggerLevel = LoggerLevel(logrus.DebugLevel)
	LevelInfo  LoggerLevel = LoggerLevel(logrus.InfoLevel)
	LevelWarn  LoggerLevel = LoggerLevel(logrus.WarnLevel)
	LevelError LoggerLevel = LoggerLevel(logrus.ErrorLevel)
)

var (
	once     sync.Once
	instance Logger
)

func NewLogger() Logger {
	once.Do(func() {
		instance = initLogger()
	})
	return instance
}

func initLogger() Logger {
	logFile := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10, // MB
		MaxBackups: 3,
		MaxAge:     30, // days
		Compress:   true,
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

type LoggerImpl struct {
	consoleEntry *logrus.Entry
}

func (l *LoggerImpl) SetLevel(level LoggerLevel) {
	l.consoleEntry.Logger.SetLevel(logrus.Level(level))
}

func (l *LoggerImpl) SetOutput(output io.Writer) {
	l.consoleEntry.Logger.SetOutput(output)
}

func (l *LoggerImpl) GetLevel() LoggerLevel {
	return LoggerLevel(l.consoleEntry.Logger.GetLevel())
}

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

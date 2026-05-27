// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package logutils

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Pre-define log instance with default info level.
var (
	Logger       *zap.Logger
	LoggerStderr *zap.Logger
)

// LogFormat is a type help you choose the log output format, which supports "json" and "console".
type LogFormat string

const (
	JSONLogFormat    LogFormat = "json"
	ConsoleLogFormat LogFormat = "console"
)

// FileOutputOption supports the configuration for log file
type FileOutputOption struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

const (
	DefaultIPAMPluginLogFilePath = "/var/log/spidernet/spiderpool.log"
	// MaxSize    = 100 // MB
	DefaultLogFileMaxSize int = 100
	// MaxAge     = 30 // days (no limit)
	DefaultLogFileMaxAge = 30
	// MaxBackups = 10 // no limit
	DefaultLogFileMaxBackups = 10
)

// LogMode is a type help you choose the log output mode, which supports "stderr","stdout","file".
type LogMode uint32

const (
	OutputFile   LogMode = 1 // 0001
	OutputStderr LogMode = 2 // 0010
	OutputStdout LogMode = 4 // 0100
)

// LogLevel is an alias for zapcore.Level.
type LogLevel = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

// Log level character string
const (
	LogDebugLevelStr = "debug"
	LogInfoLevelStr  = "info"
	LogWarnLevelStr  = "warn"
	LogErrorLevelStr = "error"
	LogFatalLevelStr = "fatal"
	LogPanicLevelStr = "panic"
)

func ConvertLogLevel(level string) *LogLevel { _ = "STUB: not implemented"; return nil }

func init() {
	err := InitStdoutLogger(InfoLevel)
	if nil != err {
		panic(err)
	}

	err = InitStderrLogger(InfoLevel)
	if nil != err {
		panic(err)
	}
}

// NewLoggerWithOption provides the ability to custom log with options.
// You can choose 'output format', 'output mode' and decide to use 'time prefix', 'function caller suffix'
// 'log level prefix' or not.
// If you choose 'file output mode', you can use 'stdout and file' or 'stderr and file' together with '|'.
// The param fileOutputOption should be a pointer for FileOutputOption , and it could be nill.
// If the param isn't nil, the FileOutputOption.MaxSize, FileOutputOption.MaxAge,
// and FileOutputOption.MaxBackups have to be nonnegative number, or they will be set to default value.
func NewLoggerWithOption(
	format LogFormat,
	outputMode LogMode,
	fileOutputOption *FileOutputOption,
	addTimePrefix, addLogLevelPrefix, addFuncCallerSuffix bool,
	logLevel LogLevel,
) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	// MaxSize    = 100 // MB
	// MaxAge     = 30 // days (no limit)
	// MaxBackups = 10 // no limit
	// LocalTime  = false // use computers local time, UTC by default
	// Compress   = false // compress the rotated log in gzip format
	return nil, nil
}

// set zap encoder configuration

// InitStdoutLogger create  Logger instance with default configuration for 'stdout' usage, it's JSONLogFormat.
func InitStdoutLogger(logLevel LogLevel) error { _ = "STUB: not implemented"; return nil }

// InitStderrLogger create LoggerStderr instance for 'stderr' usage, it's ConsoleLogFormat.
// It wouldn't provide 'time prefix', 'function caller suffix' and 'log level prefix' in output.
func InitStderrLogger(logLevel LogLevel) error { _ = "STUB: not implemented"; return nil }

// InitFileLogger sets LoggerFile configuration for 'file output' usage.
// fileMaxSize unit MB, fileMaxAge unit days, fileMaxBackups unit counts.
func InitFileLogger(logLevel LogLevel, filePath string, fileMaxSize, fileMaxAge, fileMaxBackups int) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetupFileLogging Set up file logging.
func SetupFileLogging(level, logPath string, maxSize, maxAge, maxCount int) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loggerKey is how we find Loggers in a context.Context.
type loggerKey struct{}

// FromContext returns a logger with predefined values from a context.Context.
func FromContext(ctx context.Context) *zap.Logger { _ = "STUB: not implemented"; return nil }

// IntoContext takes a context and sets the logger as one of its values.
// Use FromContext function to retrieve the logger.
func IntoContext(ctx context.Context, logger *zap.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

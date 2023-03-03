package xlogger

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *logrus.Logger
}

// NewLogger Constructor for kratos Logger
func NewLogger(opts ...Option) log.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	for _, op := range opts {
		op(logger)
	}

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) (err error) {
	var (
		logrusLevel logrus.Level
		fields      logrus.Fields = make(map[string]interface{})
		msg         string
	)

	switch level {
	case log.LevelDebug:
		logrusLevel = logrus.DebugLevel
	case log.LevelInfo:
		logrusLevel = logrus.InfoLevel
	case log.LevelWarn:
		logrusLevel = logrus.WarnLevel
	case log.LevelError:
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.DebugLevel
	}

	if logrusLevel > l.log.Level {
		return
	}

	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue
		}
		if key == logrus.FieldKeyMsg {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		fields[key] = keyvals[i+1]
	}

	if len(fields) > 0 {
		l.log.WithFields(fields).Log(logrusLevel, msg)
	} else {
		l.log.Log(logrusLevel, msg)
	}

	return
}

type Option func(log *logrus.Logger)

// Level Specify log level for logger, default to Info level.
func Level(levelStr string) Option {
	var level logrus.Level

	switch levelStr {
	case LogLevelDebug:
		level = logrus.DebugLevel
	case LogLevelInfo:
		level = logrus.InfoLevel
	case LogLevelWarn:
		level = logrus.WarnLevel
	case LogLevelError:
		level = logrus.ErrorLevel
	default:
		level = logrus.InfoLevel
	}

	return func(log *logrus.Logger) {
		log.Level = level
	}
}

func Output(w io.Writer) Option {
	return func(log *logrus.Logger) {
		log.Out = w
	}
}

func Formatter(formatter logrus.Formatter) Option {
	return func(log *logrus.Logger) {
		log.Formatter = formatter
	}
}

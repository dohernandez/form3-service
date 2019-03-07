package log

import (
	"context"

	logging "github.com/hellofresh/logging-go"
	"github.com/sirupsen/logrus"
)

// NewLog creates Logger
func NewLog(config logging.LogConfig) (*logrus.Logger, error) {
	if config.Level == "" {
		config.Level = "debug"
	}
	err := config.Apply()
	if err != nil {
		return nil, err
	}

	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}

	ll := logrus.GetLevel()
	l.SetLevel(ll)

	return l, nil
}

type contextKey string

const ctxKey = contextKey("logger")

// ToContext sets logger instance to context
func ToContext(ctx context.Context, logger *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}

// FromContext extracts logger instance from context
func FromContext(ctx context.Context) *logrus.Logger {
	value := ctx.Value(ctxKey)
	if value == nil {
		return nil
	}

	return value.(*logrus.Logger)
}

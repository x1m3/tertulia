package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

type Fields map[string]interface{}

func Debug(ctx context.Context, msg string) {
	logrus.WithContext(ctx).Debug(msg)
}

func Debugf(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).WithFields(logrus.Fields(fields)).Debug(msg)
}

func Info(ctx context.Context, msg string) {
	logrus.WithContext(ctx).Info(msg)
}

func Infof(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).WithFields(logrus.Fields(fields)).Info(msg)
}

func Warn(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).Warn(msg)
}

func Warnf(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).WithFields(logrus.Fields(fields)).Warn(msg)
}

func Error(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).Error(msg)
}

func Errorf(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).WithFields(logrus.Fields(fields)).Error(msg)
}

func Fatal(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).Fatal(msg)
}

func Fatalf(ctx context.Context, msg string, fields Fields) {
	logrus.WithContext(ctx).WithFields(logrus.Fields(fields)).Fatal(msg)
}

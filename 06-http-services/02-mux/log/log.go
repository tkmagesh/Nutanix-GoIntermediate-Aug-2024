package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type LogFields map[string]any

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Error(ctx context.Context, msg string, fields LogFields) {
	fields["request-id"] = ctx.Value("request-id")
	log.WithFields(log.Fields(fields)).Error(msg)
}

func Warning(ctx context.Context, msg string, fields LogFields) {
	fields["request-id"] = ctx.Value("request-id")
	log.WithFields(log.Fields(fields)).Warn(msg)
}

func Info(ctx context.Context, msg string, fields LogFields) {
	fields["request-id"] = ctx.Value("request-id")
	log.WithFields(log.Fields(fields)).Info(msg)
}

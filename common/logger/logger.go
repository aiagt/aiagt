package logger

import (
	ktlog "github.com/aiagt/kitextool/option/server/log"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
)

var logger *kitexzap.Logger

func init() {
	logger = ktlog.NewZapLogger()
}

func Logger() *kitexzap.Logger {
	return logger
}

func With(args ...interface{}) *kitexzap.Logger {
	t := *logger
	t.SugaredLogger = t.With(args...)

	return &t
}

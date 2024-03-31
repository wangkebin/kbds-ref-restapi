package dal

import (
	"fmt"
	"time"

	zl "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gl "gorm.io/gorm/logger"
)

type logwritter struct {
	loglevel gl.LogLevel
	logger   *zl.Logger
}

func (lw *logwritter) Printf(template string, data ...interface{}) {

	zaps := make([]zapcore.Field, 0)
	for d := range data {
		zaps = append(zaps, zl.String("", fmt.Sprint(d)))
	}
	switch lw.loglevel {
	case gl.Error:
		lw.logger.Error(template, zaps...)
	case gl.Info:
		lw.logger.Info(template, zaps...)
	case gl.Warn:
		lw.logger.Warn(template, zaps...)
	case gl.Silent:
		{
		}
	default:
		lw.logger.Info(template, zaps...)
	}
}

func LogConfig(l *zl.Logger, lv gl.LogLevel) *gl.Interface {
	var lw = logwritter{
		logger:   l,
		loglevel: lv,
	}

	var cfg = gl.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  lw.loglevel,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}
	var lgi = gl.New(&lw, cfg)

	return &lgi
}

package logger_test

import (
	"fmt"
	"testing"

	"github.com/roman-kartash/gocommons/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(t *testing.T, lvl string, dev bool) *zap.Logger {
	t.Helper()

	cfg := &logger.Config{
		Lvl:        lvl,
		File:       fmt.Sprintf("../tmp/main_%s.log", lvl),
		Dev:        dev,
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     90,
		Compress:   false,
	}

	require.NoError(t, cfg.AfterLoad())

	l, syncL, err := logger.NewLogger(cfg)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = syncL()
	})

	return l
}

func GetConsoleLogger(t *testing.T, lvl zapcore.Level) *zap.Logger {
	t.Helper()

	l, syncL := logger.NewConsoleLogger(lvl, true)

	t.Cleanup(func() {
		_ = syncL()
	})

	return l
}

func TestLoggers(t *testing.T) {
	t.Parallel()

	loggers := map[string]*zap.Logger{
		"debug":   GetLogger(t, logger.DebugLevelStr, true),
		"info":    GetLogger(t, logger.InfoLevelStr, false),
		"warning": GetLogger(t, logger.WarningLevelStr, false),
		"error":   GetLogger(t, logger.ErrorLevelStr, false),

		"console debug":   GetConsoleLogger(t, zapcore.DebugLevel),
		"console info":    GetConsoleLogger(t, zapcore.InfoLevel),
		"console warning": GetConsoleLogger(t, zapcore.WarnLevel),
		"console error":   GetConsoleLogger(t, zapcore.ErrorLevel),
	}

	for n, l := range loggers {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			l = l.Named(t.Name())

			l.Debug("debug message", zap.Int("int", 10), zap.String("str", "test"))
			l.Info("info message", zap.Int("int", 10), zap.String("str", "test"))
			l.Warn("warn message", zap.Int("int", 10), zap.String("str", "test"))
			l.Error("error message", zap.Int("int", 10), zap.String("str", "test"))

			require.Panics(t, func() {
				l.Panic("panic happens")
			})
		})
	}
}

package postgres

import (
	"context"
	"log/slog"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/tracelog"
)

type Logger struct {
	l *slog.Logger
}

func NewTracer(l *slog.Logger) pgx.QueryTracer {
	return &tracelog.TraceLog{
		Logger:   &Logger{l: l},
		LogLevel: tracelog.LogLevelTrace,
	}
}

func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	logger := l.l
	attrs := make([]slog.Attr, 0, len(data))
	keys := sortedKeys(data)
	for _, k := range keys {
		v := data[k]
		if k == "sql" {
			str, ok := v.(string)
			if ok {
				v = PrettySQL(str)
			}
		}
		attrs = append(attrs, slog.Any(k, v))
	}

	logger.LogAttrs(ctx, translateLevel(level), msg, attrs...)
}

func translateLevel(level tracelog.LogLevel) slog.Level {
	switch level {
	case tracelog.LogLevelTrace:
		return slog.LevelDebug
	case tracelog.LogLevelDebug:
		return slog.LevelDebug
	case tracelog.LogLevelInfo:
		return slog.LevelInfo
	case tracelog.LogLevelWarn:
		return slog.LevelWarn
	case tracelog.LogLevelError:
		return slog.LevelError
	case tracelog.LogLevelNone:
		return slog.LevelError
	default:
		return slog.LevelError
	}
}

func sortedKeys(data map[string]any) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

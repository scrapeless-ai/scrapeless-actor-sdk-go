package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type TraceIDHook struct{}

func (hook *TraceIDHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *TraceIDHook) Fire(entry *log.Entry) error {
	if entry.Context != nil {
		traceID := entry.Context.Value("trace-id")
		if traceID != nil {
			entry.Data["trace-id"] = traceID
		}
	}

	return nil
}

func InitLogrus(levelStr string) {
	level := log.TraceLevel
	if l, ok := LogLevel[strings.ToLower(levelStr)]; ok {
		level = l
	}

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
		ForceColors:   true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			fc := path.Base(f.Function)
			return fmt.Sprintf("%s()", fc), fmt.Sprintf(" - %s:%d", filename, f.Line)
		},
		TimestampFormat: time.DateTime,
	})
	log.SetReportCaller(true)
	log.SetLevel(level)

	log.AddHook(&TraceIDHook{})
}

var LogLevel = map[string]log.Level{
	"panic": log.PanicLevel,
	"fatal": log.FatalLevel,
	"error": log.ErrorLevel,
	"warn":  log.WarnLevel,
	"info":  log.InfoLevel,
	"debug": log.DebugLevel,
	"trace": log.TraceLevel,
}

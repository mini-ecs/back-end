package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

var _log *logrus.Logger

func GetZtFormatter() logrus.Formatter {
	var exampleFormatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			HideKeys:     true,
			FieldsOrder:  []string{"component", "category"},
			TrimMessages: true,
		},
	}
	return exampleFormatter
}

func GetNestedFormatter() logrus.Formatter {
	return &nested.Formatter{
		HideKeys:     true,
		FieldsOrder:  []string{"component", "category"},
		TrimMessages: true,
		NoColors:     false,
	}
}
func GetDefaultFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			funcName := strings.Split(frame.Function, ".")
			return funcName[len(funcName)-1] + " :", fileName
		},
	}
}

func SetLogToFile(l *logrus.Logger) {
	file, err := os.OpenFile(config.GetConfig().Log.Path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		l.Out = file
	} else {
		l.Info("Failed to log to file, using default stderr")
	}
}

func switchLevel(l string) logrus.Level {
	switch l {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.DebugLevel

	}
}

func GetGlobalLogger() *logrus.Logger {
	once := sync.Once{}
	once.Do(func() {
		_log = logrus.StandardLogger()
		_log.SetReportCaller(true)
		_log.SetFormatter(GetZtFormatter())
		_log.SetLevel(switchLevel(config.GetConfig().Log.Level))
	})
	return _log
}

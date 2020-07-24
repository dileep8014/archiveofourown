package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	pool sync.Pool
	once sync.Once
)

func SetupLogger() {
	once.Do(func() {
		if global.ServerSetting.RunMode == "debug" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		consoleWriter := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.RFC3339
			w.Out = io.MultiWriter(os.Stdout, &lumberjack.Logger{
				Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
				MaxSize:  600,
				MaxAge:   10,
			})
			w.FormatLevel = func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("| %-4s ", i))
			}
			w.FormatMessage = func(i interface{}) string {
				if i != nil {
					return fmt.Sprintf("| ***%s****", i)
				}
				return ""
			}
			w.FormatErrFieldName = func(i interface{}) string {
				return fmt.Sprintf("| %s:", i)
			}
			w.FormatErrFieldValue = func(i interface{}) string {
				return fmt.Sprintf("%s ", i)
			}
			w.FormatFieldName = w.FormatErrFieldName
			w.FormatFieldValue = w.FormatErrFieldValue
			w.FormatTimestamp = func(i interface{}) string {
				return fmt.Sprintf("%s ", i)
			}
			w.FormatCaller = func(i interface{}) string {
				return fmt.Sprintf("| %s", i)
			}
		})
		pool.New = func() interface{} {
			return zerolog.New(consoleWriter).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		}
	})
}

func Get() zerolog.Logger {
	v := pool.Get()
	if v == nil {
		return zerolog.Nop()
	}
	return v.(zerolog.Logger)
}

func Put(logger zerolog.Logger) {
	pool.Put(logger)
}

package global

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

var Logger zerolog.Logger

func SetupLogger() {

	if ServerSetting.RunMode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	consoleWriter := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
		w.Out = io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt,
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
			if i != nil {
				return fmt.Sprintf("| %s", i)
			}
			return ""
		}
	})

	Logger = zerolog.New(consoleWriter).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

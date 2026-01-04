package zl

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ansufw/gomono/internal/shared/enum"
	"github.com/rs/zerolog"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorGray   = "\033[90m"
)

func Logger(m enum.Mode, s enum.Service) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout}

	if m == enum.ModeDev {

		output.FormatTimestamp = func(i any) string {
			if ts, ok := i.(string); ok {
				t, err := time.Parse(time.RFC3339, ts)
				if err == nil {
					return colorGray + t.Format("02/01/06T15:04:05Z07:00 |") + colorReset
				}
			}
			return fmt.Sprintf("%s |", i)
		}

		output.FormatMessage = func(i any) string {
			return fmt.Sprintf("%s", i)
		}

		output.FormatLevel = func(i any) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}

		output.FormatFieldName = func(i any) string {
			return fmt.Sprintf("%s:", i)
		}

		// output.FieldsOrder = []string{"service"}
		output.FormatFieldValue = func(i any) string {
			return fmt.Sprintf("%s", i)
		}

		output.PartsOrder = []string{
			zerolog.TimestampFieldName,
			"service",
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		}

	}

	logger := zerolog.New(output).
		With().Str("service", string(s)).
		Timestamp().Logger()
	return &logger
}

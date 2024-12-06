package logger

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: false,
		ReportTimestamp: true,
		TimeFormat: "2006/01/02 15:04",
		Prefix: "Gin-gorm Basic 🍪 ",

	})

	styles := GetStyles()
	logger.SetStyles(styles)
}

func Debug(message string, args ...interface{}) {

	for i := 0; i < len(args); i++ {
		// Memeriksa jika index ganjil dan bertipe string
		if i%2 != 0 {
			if str, ok := args[i].(string); ok {
				// Menambahkan simbol " di awal dan akhir string
				args[i] = fmt.Sprintf("'%s'", str)
			}
		}
	}

	logger.Debug(message, args...)
}
func Info(message string, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		// Memeriksa jika index ganjil dan bertipe string
		if i%2 != 0 {
			if str, ok := args[i].(string); ok {
				// Menambahkan simbol " di awal dan akhir string
				args[i] = fmt.Sprintf("'%s'", str)
			}
		}
	}
	logger.Info(message, args...)
}
func Error(message string, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		// Memeriksa jika index ganjil dan bertipe string
		if i%2 != 0 {
			if str, ok := args[i].(string); ok {
				// Menambahkan simbol " di awal dan akhir string
				args[i] = fmt.Sprintf("'%s'", str)
			}
		}
	}
	logger.Error(message, args...)
}
func Warn(message string, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		// Memeriksa jika index ganjil dan bertipe string
		if i%2 != 0 {
			if str, ok := args[i].(string); ok {
				// Menambahkan simbol " di awal dan akhir string
				args[i] = fmt.Sprintf("'%s'", str)
			}
		}
	}
	logger.Warn(message, args...)
}
func Fatal(message string, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		// Memeriksa jika index ganjil dan bertipe string
		if i%2 != 0 {
			if str, ok := args[i].(string); ok {
				// Menambahkan simbol " di awal dan akhir string
				args[i] = fmt.Sprintf("'%s'", str)
			}
		}
	}
	logger.Fatal(message, args...)
}
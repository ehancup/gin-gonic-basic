package logger

import (
	// "os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func GetStyles() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("[DEBU]").
		Padding(0, 1, 0, 1).
		// Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("#8b8bff")).Bold(true)

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("[INFO]").
		Padding(0, 1, 0, 1).
		// Background(lipgloss.Color("#58e5c2")).
		Foreground(lipgloss.Color("#58e5c2")).Bold(true)

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("[WARN]").
		Padding(0, 1, 0, 1).
		// Background(lipgloss.Color("#ccf281")).
		Foreground(lipgloss.Color("#ccf281")).Bold(true)

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("[ERROR]").
		Padding(0, 1, 0, 1).
		// Background(lipgloss.Color("#ff5f87")).
		Foreground(lipgloss.Color("#ff5f87")).Bold(true)

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("[FATAL]").
		Padding(0, 1, 0, 1).
		// Background(lipgloss.Color("#bf7edf")).
		Foreground(lipgloss.Color("#bf7edf")).Bold(true)

	styles.Message = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#ffffff"))
	styles.Value = lipgloss.NewStyle().Foreground(lipgloss.Color("#ababab"))
	

	return styles
}

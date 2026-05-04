package tui

import (
	"io"

	"github.com/charmbracelet/lipgloss"
)

// BindRenderer wires lipgloss to the same writer Bubble Tea uses so color/profile
// detection matches the active terminal (fewer mismatched backgrounds across emulators).
func BindRenderer(out io.Writer) {
	if out == nil {
		return
	}

	lipgloss.SetDefaultRenderer(lipgloss.NewRenderer(out))
}

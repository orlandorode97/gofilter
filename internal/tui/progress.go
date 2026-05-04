package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// viewProcessing renders the progress bar view during processing.
func (m Model) viewProcessing() string {
	sub := lipgloss.JoinVertical(
		lipgloss.Left,
		_metaAccentStyle.Render(fmt.Sprintf("Recipe · %s", m.selectedFilter)),
		_dimStyle.Render("Crunching pixels — shouldn’t be long."),
		"",
		m.progress.View(),
	)

	return _flowShell(
		m.layoutW,
		m.height,
		4,
		"Working magic",
		"You can exhale — we’ll autosave when this hits 100%.",
		sub,
		_dimStyle.Render("Tips · grab coffee · don’t resize tiny terminals mid-flight"),
	)
}

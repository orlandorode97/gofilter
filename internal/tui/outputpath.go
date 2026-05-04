package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// viewOutputPath renders the output path configuration view.
func (m Model) viewOutputPath() string {
	meta := lipgloss.JoinVertical(
		lipgloss.Left,
		_pathSummary(m.layoutW, m.selectedImg),
		"",
		_metaAccentStyle.Render("Active filter · "+m.selectedFilter),
	)

	field := lipgloss.JoinVertical(
		lipgloss.Left,
		_bodyStyle.Bold(true).Render("Drop files here (folder path)"),
		m.outputPath.View(),
	)

	inner := lipgloss.JoinVertical(lipgloss.Left, meta, "", field)

	return _flowShell(
		m.layoutW,
		m.height,
		3,
		"Pick an output home",
		"Must be a folder you can write to — we’ll drop the finished image inside.",
		inner,
		_footerOutput(),
	)
}

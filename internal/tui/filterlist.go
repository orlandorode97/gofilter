package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// viewFilterList renders the filter selection view.
func (m Model) viewFilterList() string {
	descW := m.layoutW - 10
	if descW < 28 {
		descW = 28
	}

	var sb strings.Builder

	for i, f := range m.filters {
		if i > 0 {
			sb.WriteByte('\n')
		}

		if i == m.cursor {
			sb.WriteString(_selLineStyle.Width(descW).Render("› " + f.Name))
			sb.WriteByte('\n')
			sb.WriteString(_selDescStyle.Width(descW).Render(f.Description))

			continue
		}

		sb.WriteString(_rowStyle.Width(descW).Render("  " + f.Name))
		sb.WriteByte('\n')
		sb.WriteString(_dimStyle.Width(descW).Render("  " + f.Description))
	}

	summary := _pathSummary(m.layoutW, m.selectedImg)

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		sb.String(),
		"",
		summary,
	)

	return _flowShell(
		m.layoutW,
		m.height,
		2,
		"Dial in the vibe",
		"Highlight a row — description updates with your cursor.",
		inner,
		_footerFilter(),
	)
}

package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// viewSuccess renders the success message view.
func (m Model) viewSuccess() string {
	spark := lipgloss.NewStyle().Foreground(lipgloss.Color(_hexGold)).Bold(true).Render(" ★ ")
	head := lipgloss.JoinHorizontal(lipgloss.Left,
		spark,
		_okTitleStyle.Render("Nailed it."),
		spark,
	)

	hintW := m.layoutW - 10
	if hintW < 28 {
		hintW = 28
	}

	body := _boxedBody(
		m.layoutW,
		lipgloss.JoinVertical(
			lipgloss.Left,
			_metaAccentStyle.Render("Fresh file"),
			_pathSummary(m.layoutW, m.finalOutputPath),
			"",
		),
	)

	airGap := ""
	switch {
	case m.height >= 46:
		airGap = "\n\n"
	case m.height >= 34:
		airGap = "\n"
	default:
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		_brandHeader(),
		airGap,
		head,
		"",
		body,
		airGap,
		_footerDismiss(),
	)
}

// viewError renders the error message view.
func (m Model) viewError() string {
	head := lipgloss.JoinHorizontal(lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color(_hexRose)).Bold(true).Render(" ▸ "),
		_errTitleStyle.Render("That didn’t fly"),
	)

	errW := m.layoutW - 8
	if errW < 28 {
		errW = 28
	}

	body := _boxedBody(
		m.layoutW,
		lipgloss.JoinVertical(
			lipgloss.Left,
			_bodyStyle.Width(errW).Render(m.err.Error()),
			"",
			_dimStyle.Width(errW).Render("Fix the issue and launch gofilter again — your picks aren’t saved."),
		),
	)

	airGap := ""
	switch {
	case m.height >= 46:
		airGap = "\n\n"
	case m.height >= 34:
		airGap = "\n"
	default:
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		_brandHeader(),
		airGap,
		head,
		"",
		body,
		airGap,
		_footerDismiss(),
	)
}

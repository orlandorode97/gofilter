package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	_flowStepsTotal = 4

	// Responsive layout: content tracks terminal columns within these bounds.
	_layoutSideMarginTotal = 10 // subtract from term width (margins + slack)
	_layoutMinWidth        = 52
	_layoutMaxWidth        = 180
	_layoutDefaultWidth    = 92 // until the first WindowSizeMsg arrives
)

// Palette — neon accents on the terminal’s default background (no ANSI fills).
// Background colors composite inconsistently across Terminal.app / Ghostty / WezTerm.
const (
	_hexNeonMagenta = "#FF2D92"
	_hexNeonCyan    = "#00F5FF"
	_hexGold        = "#FFD60A"
	_hexPurpleGlow  = "#A855F7"
	_hexText        = "#F1F5F9"
	_hexMuted       = "#8B9CB3"
	_hexRow         = "#CBD5E1"
	_hexFile        = "#67E8F9"
	_hexMint        = "#34F5C5"
	_hexRose        = "#FF4D6D"
	_hexDescHi      = "#E9D5FF"
)

var (
	_stepTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexNeonCyan)).
			Bold(true)

	_taglineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexMuted))

	_bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexText))

	_dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexMuted))

	_selLineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexGold)).
			Bold(true)

	_selDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexDescHi))

	_rowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexRow))

	_borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(_hexNeonMagenta)).
			Padding(0, 1)

	_metaAccentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(_hexPurpleGlow))

	_baseFileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(_hexFile)).
			Bold(true)

	_errTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(_hexRose))

	_okTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(_hexMint))
)

// layoutContentWidth maps terminal columns to the bordered panel width.
func layoutContentWidth(termCols int) int {
	if termCols <= 0 {
		return _layoutDefaultWidth
	}

	w := termCols - _layoutSideMarginTotal
	if w < _layoutMinWidth {
		return _layoutMinWidth
	}

	if w > _layoutMaxWidth {
		return _layoutMaxWidth
	}

	return w
}

// progressBarViewportWidth maps layout width to bubbles progress width.
func progressBarViewportWidth(layoutW int) int {
	w := layoutW - 12
	if w < 28 {
		return 28
	}

	return w
}

// textInputViewportWidth maps layout width to bubbles textinput viewport width.
func textInputViewportWidth(layoutW int) int {
	w := layoutW - 14
	if w < 28 {
		return 28
	}

	return w
}

// _brandHeader is the app masthead — quick read, high energy.
func _brandHeader() string {
	wordmark := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(_hexGold)).
		Render("gofilter")

	flair := lipgloss.NewStyle().
		Foreground(lipgloss.Color(_hexNeonMagenta)).
		Bold(true).
		Render("✦")

	return lipgloss.JoinHorizontal(lipgloss.Left, wordmark, " ", flair)
}

// _stepRail shows journey progress across four beats.
func _stepRail(activeStep int) string {
	var b strings.Builder

	for i := 1; i <= _flowStepsTotal; i++ {
		if i > 1 {
			if activeStep >= i {
				b.WriteString(lipgloss.NewStyle().
					Foreground(lipgloss.Color(_hexGold)).
					Render(" ─── "))
			} else {
				b.WriteString(_dimStyle.Render("  ·  "))
			}
		}

		if activeStep >= i {
			b.WriteString(lipgloss.NewStyle().
				Foreground(lipgloss.Color(_hexNeonMagenta)).
				Bold(true).
				Render("●"))
			continue
		}

		b.WriteString(_dimStyle.Render("○"))
	}

	return b.String()
}

func _kbdLabel(keys, caption string) string {
	muted := lipgloss.NewStyle().Foreground(lipgloss.Color(_hexMuted))

	keyStyled := lipgloss.NewStyle().
		Foreground(lipgloss.Color(_hexNeonCyan)).
		Bold(true).
		Render(keys)

	chips := muted.Render("[") + keyStyled + muted.Render("]")

	return lipgloss.JoinHorizontal(lipgloss.Left,
		chips,
		muted.Render(" "+caption+"   "),
	)
}

func _footerBrowse() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		_kbdLabel("↑↓", "move"),
		_kbdLabel("enter", "open / pick"),
		_kbdLabel("q", "quit"),
	)
}

func _footerFilter() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		_kbdLabel("↑↓", "move"),
		_kbdLabel("enter", "choose"),
		_kbdLabel("q", "quit"),
	)
}

func _footerOutput() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		_kbdLabel("enter", "run"),
		_kbdLabel("esc", "back"),
		_kbdLabel("q", "quit"),
	)
}

func _footerDismiss() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		_kbdLabel("any", "exit"),
	)
}

// _flowShell wraps main wizard screens with shared chrome + UX rhythm.
func _flowShell(contentW, termH int, step int, heading, tagline, innerBody, footer string) string {
	airGap := ""
	switch {
	case termH >= 46:
		airGap = "\n\n"
	case termH >= 34:
		airGap = "\n"
	default:
	}

	block := lipgloss.JoinVertical(
		lipgloss.Left,
		_brandHeader(),
		airGap,
		_stepRail(step),
		"",
		_stepTitleStyle.Render(heading),
		_taglineStyle.Render(tagline),
		"",
		_boxedBody(contentW, innerBody),
		airGap,
		footer,
	)

	return block
}

func _boxedBody(contentW int, body string) string {
	return _borderStyle.Width(contentW).Render(strings.TrimRight(body, "\n"))
}

func _pathSummary(layoutW int, fullPath string) string {
	if fullPath == "" {
		return ""
	}

	base := filepath.Base(fullPath)
	dir := filepath.Dir(fullPath)

	textW := layoutW - 8
	if textW < 28 {
		textW = 28
	}

	baseLine := _baseFileStyle.Width(textW).Render(base)
	dirLine := _dimStyle.Width(textW).Render(dir)

	return fmt.Sprintf("%s\n%s", baseLine, dirLine)
}

func _backdropPlace(termW, termH int, body string) string {
	if termW <= 0 || termH <= 0 {
		return body
	}

	// No whitespace background fill — avoids solid tint rectangles that diverge
	// between Terminal.app, Ghostty, and WezTerm when layered with bordered panels.
	return lipgloss.Place(termW, termH, lipgloss.Center, lipgloss.Center, body)
}

package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Custom file picker that works reliably.
type CustomFilePicker struct {
	CurrentDirectory string
	Files            []string
	SelectedIndex    int
	ShowHidden       bool
}

// NewCustomFilePicker creates a new file picker.
func NewCustomFilePicker(startDir string) *CustomFilePicker {
	fp := &CustomFilePicker{
		CurrentDirectory: startDir,
		ShowHidden:       false,
	}
	fp.ReadDir()

	return fp
}

// ReadDir reads the current directory.
func (fp *CustomFilePicker) ReadDir() {
	entries, err := os.ReadDir(fp.CurrentDirectory)
	if err != nil {
		return
	}

	fp.Files = []string{}

	// Add parent directory.
	fp.Files = append(fp.Files, "..")

	// Sort: directories first, then files.
	type fileEntry struct {
		name  string
		isDir bool
	}

	var dirs, files []fileEntry
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") && !fp.ShowHidden {
			continue
		}

		ext := strings.ToLower(filepath.Ext(e.Name()))
		if !e.IsDir() && ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			continue
		}

		if e.IsDir() {
			dirs = append(dirs, fileEntry{name: e.Name(), isDir: true})
		} else {
			files = append(files, fileEntry{name: e.Name(), isDir: false})
		}
	}

	sort.Slice(dirs, func(i, j int) bool { return dirs[i].name < dirs[j].name })
	sort.Slice(files, func(i, j int) bool { return files[i].name < files[j].name })

	for _, d := range dirs {
		fp.Files = append(fp.Files, d.name+"/")
	}
	for _, f := range files {
		fp.Files = append(fp.Files, f.name)
	}

	fp.SelectedIndex = 0
}

// View renders the file picker with wrapping tuned to panel width (columns).
func (fp *CustomFilePicker) View(contentW int) string {
	textW := contentW - 10
	if textW < 28 {
		textW = 28
	}

	var sb strings.Builder

	for i, f := range fp.Files {
		line := "› " + f

		switch {
		case i == fp.SelectedIndex:
			sb.WriteString(_selLineStyle.Width(textW).Render(line))
		default:
			sb.WriteString(_rowStyle.Width(textW).Render("  " + f))
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

// MoveUp moves the cursor up.
func (fp *CustomFilePicker) MoveUp() {
	if fp.SelectedIndex > 0 {
		fp.SelectedIndex--
	}
}

// MoveDown moves the cursor down.
func (fp *CustomFilePicker) MoveDown() {
	if fp.SelectedIndex < len(fp.Files)-1 {
		fp.SelectedIndex++
	}
}

// Select selects the current file or directory.
// Returns (selectedPath, isDir, error).
func (fp *CustomFilePicker) Select() (string, bool, error) {
	if len(fp.Files) == 0 {
		return "", false, fmt.Errorf("no file selected")
	}

	selected := fp.Files[fp.SelectedIndex]

	// Handle parent directory.
	if selected == "../" || selected == ".." {
		parent := filepath.Dir(fp.CurrentDirectory)
		fp.CurrentDirectory = parent
		fp.ReadDir()

		return "", true, nil
	}

	fullPath := filepath.Join(fp.CurrentDirectory, strings.TrimSuffix(selected, "/"))
	info, err := os.Stat(fullPath)
	if err != nil {
		return "", false, err
	}

	if info.IsDir() {
		fp.CurrentDirectory = fullPath
		fp.ReadDir()

		return "", true, nil
	}

	return fullPath, false, nil
}

// viewFilePicker renders the custom file picker view.
func (m Model) viewFilePicker() string {
	dirW := m.layoutW - 10
	if dirW < 28 {
		dirW = 28
	}

	dirLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color(_hexPurpleGlow)).
		Width(dirW).
		Render(m.customFP.CurrentDirectory)

	listView := m.customFP.View(m.layoutW)
	if len(m.customFP.Files) == 1 && m.customFP.Files[0] == ".." {
		listView += _dimStyle.Render("\n(no JPEG or PNG here — open another folder)")
	}

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		dirLine,
		"",
		listView,
	)

	return _flowShell(
		m.layoutW,
		m.height,
		1,
		"Choose your image",
		"We keep it simple: JPEG & PNG only.",
		inner,
		_footerBrowse(),
	)
}

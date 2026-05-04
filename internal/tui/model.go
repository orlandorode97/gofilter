package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/OrlandoRomo/go-filter/internal/filter"
	"github.com/OrlandoRomo/go-filter/internal/image"
)

// Progress constants.
const (
	_progressIncrement = 0.05
	_progressTickDelay = 200 * time.Millisecond
	_progressMax       = 0.85
	_progressDone      = 1.0
)

// State represents the current screen/state of the TUI.
type State int

const (
	StateFilePicker State = iota
	StateFilterList
	StateOutputPath
	StateProcessing
	StateSuccess
	StateError
)

// Model is the main Bubble Tea model for the application.
type Model struct {
	state   State
	width   int
	height  int
	layoutW int // bordered panel width derived from terminal columns

	// Custom file picker
	customFP    *CustomFilePicker
	selectedImg string

	// Filter list
	filters        []filter.FilterInfo
	cursor         int
	selectedFilter string

	// Output path
	outputPath      textinput.Model
	defaultPath     string
	finalOutputPath string

	// Progress
	progress progress.Model
	err      error

	// When true, image work finished and we stay on the processing view until the bar finishes animating to 100%.
	processingComplete bool

	// Processing
	progressChan chan float64
	doneChan     chan string
}

// NewModel creates a new TUI model with default values.
func NewModel() Model {
	homeDir, err := os.UserHomeDir()
	if err != nil || homeDir == "" {
		homeDir = "."
	}

	layoutW := layoutContentWidth(0)

	ti := textinput.New()
	ti.SetValue(homeDir)
	ti.Focus()
	ti.Width = textInputViewportWidth(layoutW)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(_hexPurpleGlow))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(_hexText))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(_hexMuted))
	ti.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(_hexGold)).
		Bold(true).
		Reverse(true)

	p := progress.New(
		progress.WithSolidFill(_hexNeonCyan),
		progress.WithWidth(progressBarViewportWidth(layoutW)),
	)

	return Model{
		state:        StateFilePicker,
		customFP:     NewCustomFilePicker(homeDir),
		filters:      filter.GetFilterList(),
		outputPath:   ti,
		defaultPath:  homeDir,
		layoutW:      layoutW,
		progress:     p,
		progressChan: make(chan float64),
		doneChan:     make(chan string),
	}
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
	)
}

// Update handles messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layoutW = layoutContentWidth(msg.Width)
		m.progress.Width = progressBarViewportWidth(m.layoutW)
		m.outputPath.Width = textInputViewportWidth(m.layoutW)

		return m, nil

	case tea.KeyMsg:
		switch m.state {
		case StateFilePicker:
			return m.updateFilePicker(msg)
		case StateFilterList:
			return m.updateFilterList(msg)
		case StateOutputPath:
			return m.updateOutputPath(msg)
		case StateProcessing:
			return m, nil
		case StateSuccess:
			return m.updateSuccess(msg)
		case StateError:
			return m.updateError(msg)
		}

	case progress.FrameMsg:
		pm, cmd := m.progress.Update(msg)
		next, ok := pm.(progress.Model)
		if !ok {
			return m, cmd
		}
		m.progress = next
		if m.state == StateProcessing && m.processingComplete && !m.progress.IsAnimating() {
			m.state = StateSuccess
			m.processingComplete = false
		}

		return m, cmd

	case progressTickMsg:
		current := m.progress.Percent()
		if current < _progressMax {
			cmd := m.progress.IncrPercent(_progressIncrement)
			return m, tea.Batch(cmd, tea.Tick(_progressTickDelay, func(t time.Time) tea.Msg {
				return progressTickMsg{}
			}))
		}

		return m, nil

	case doneMsg:
		m.finalOutputPath = string(msg)
		m.processingComplete = true
		cmd := m.progress.SetPercent(_progressDone)

		return m, cmd

	case errMsg:
		m.err = msg
		m.state = StateError

		return m, nil
	}

	return m, nil
}

// View renders the current view based on state.
func (m Model) View() string {
	var body string

	switch m.state {
	case StateFilePicker:
		body = m.viewFilePicker()
	case StateFilterList:
		body = m.viewFilterList()
	case StateOutputPath:
		body = m.viewOutputPath()
	case StateProcessing:
		body = m.viewProcessing()
	case StateSuccess:
		body = m.viewSuccess()
	case StateError:
		body = m.viewError()
	default:
		body = "Unknown state"
	}

	return _backdropPlace(m.width, m.height, body)
}

// updateFilePicker handles file picker updates.
func (m Model) updateFilePicker(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "up", "k":
		m.customFP.MoveUp()

		return m, nil

	case "down", "j":
		m.customFP.MoveDown()

		return m, nil

	case "enter":
		path, isDir, err := m.customFP.Select()
		if err != nil {
			m.err = err
			m.state = StateError

			return m, nil
		}

		// If it was a directory, just refresh the view.
		if isDir {
			return m, nil
		}

		// Validate file extension.
		ext := filepath.Ext(path)
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			m.err = fmt.Errorf("unsupported file format: %s (only .png, .jpg, .jpeg are supported)", ext)
			m.state = StateError

			return m, nil
		}

		m.selectedImg = path
		m.state = StateFilterList

		return m, nil
	}

	return m, nil
}

// updateFilterList handles filter list updates.
func (m Model) updateFilterList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.filters)-1 {
			m.cursor++
		}

	case "enter":
		m.selectedFilter = m.filters[m.cursor].Name
		m.state = StateOutputPath

		return m, nil
	}

	return m, nil
}

// updateOutputPath handles output path updates.
func (m Model) updateOutputPath(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "enter":
		outputDir := m.outputPath.Value()
		if err := image.ValidateOutputPath(outputDir); err != nil {
			m.err = err
			m.state = StateError

			return m, nil
		}

		m.state = StateProcessing
		m.processingComplete = false
		resetCmd := m.progress.SetPercent(0)

		// Start progress animation and processing in parallel.
		return m, tea.Batch(
			resetCmd,
			tea.Tick(_progressTickDelay, func(t time.Time) tea.Msg {
				return progressTickMsg{}
			}),
			m.startProcessing(outputDir),
		)

	case "esc":
		m.state = StateFilterList

		return m, nil
	}

	var cmd tea.Cmd
	m.outputPath, cmd = m.outputPath.Update(msg)

	return m, cmd
}

// updateSuccess handles success screen updates.
func (m Model) updateSuccess(_ tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

// updateError handles error screen updates.
func (m Model) updateError(_ tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

// startProcessing starts the image processing in a goroutine.
func (m Model) startProcessing(outputDir string) tea.Cmd {
	return func() tea.Msg {
		filterFunc, ok := filter.GetFilterByName(m.selectedFilter)
		if !ok {
			return errMsg(fmt.Errorf("filter %q not found", m.selectedFilter))
		}

		outputPath, err := image.ProcessAndSave(m.selectedImg, outputDir, filterFunc, nil)
		if err != nil {
			return errMsg(err)
		}

		return doneMsg(outputPath)
	}
}

// listenForProgress starts the progress ticker.
func (m Model) listenForProgress() tea.Cmd {
	return tea.Tick(_progressTickDelay, func(t time.Time) tea.Msg {
		return progressTickMsg{}
	})
}

// Custom message types.
type progressTickMsg struct{}
type doneMsg string
type errMsg error

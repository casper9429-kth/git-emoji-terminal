package gitemoji

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Emoji represents a Git emoji with its name and symbol
type Emoji struct {
	Name   string
	Symbol string
}

// List of common Git emojis
var GitEmojis = []Emoji{
	{Name: "feat", Symbol: "✨"},
	{Name: "fix", Symbol: "🐛"},
	{Name: "docs", Symbol: "📚"},
	{Name: "style", Symbol: "💎"},
	{Name: "refactor", Symbol: "📦"},
	{Name: "perf", Symbol: "🚀"},
	{Name: "test", Symbol: "🔍"},
	{Name: "build", Symbol: "🛠️"},
	{Name: "ci", Symbol: "⚙️"},
	{Name: "chore", Symbol: "♻️"},
	{Name: "revert", Symbol: "⏪"},
	{Name: "wip", Symbol: "🚧"},
	{Name: "init", Symbol: "🎉"},
	{Name: "security", Symbol: "🔒"},
	{Name: "deps", Symbol: "📦"},
	{Name: "i18n", Symbol: "🌐"},
	{Name: "typo", Symbol: "✏️"},
	{Name: "merge", Symbol: "🔀"},
	{Name: "config", Symbol: "⚙️"},
}

// Item implements the list.Item interface for our emoji list
type Item struct {
	emoji Emoji
}

func (i Item) FilterValue() string {
	return i.emoji.Name + " " + i.emoji.Symbol
}

func (i Item) Title() string {
	return i.emoji.Symbol + " " + i.emoji.Name
}

func (i Item) Description() string {
	return "Press Enter to select"
}

// EmojiPickerModel is the model for the emoji picker TUI
type EmojiPickerModel struct {
	list          list.Model
	quitting      bool
	selectedEmoji string
}

func initialModel() EmojiPickerModel {
	// Setup the emoji list items
	var items []list.Item
	for _, emoji := range GitEmojis {
		items = append(items, Item{emoji: emoji})
	}

	// Setup the list with initial size
	l := list.New(items, list.NewDefaultDelegate(), 20, 10)
	l.Title = "Select Git Emoji"
	l.SetShowFilter(true)
	l.FilterInput.Placeholder = "Type to filter emojis..."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#3C3C3C")).
		Padding(0, 1)

	return EmojiPickerModel{
		list: l,
	}
}

func (m EmojiPickerModel) Init() tea.Cmd {
	return nil
}

func (m EmojiPickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Margin(1, 2).GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// First handle any key shortcuts
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			if i, ok := m.list.SelectedItem().(Item); ok {
				m.selectedEmoji = i.emoji.Symbol
				m.quitting = true

				// Copy to clipboard
				clipboard.WriteAll(m.selectedEmoji)

				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m EmojiPickerModel) View() string {
	if m.quitting {
		if m.selectedEmoji != "" {
			return fmt.Sprintf("Selected emoji: %s (copied to clipboard)\n", m.selectedEmoji)
		}
		return "No emoji selected\n"
	}

	return m.list.View()
}

// RunPicker starts the emoji picker UI
func RunPicker() (string, error) {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run program: %v", err)
	}

	if m, ok := m.(EmojiPickerModel); ok && m.selectedEmoji != "" {
		// Try to copy to clipboard
		if err := clipboard.WriteAll(m.selectedEmoji); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not copy to clipboard: %v\n", err)
		}
		return m.selectedEmoji, nil
	}
	return "", nil
}

// ShowPicker displays the emoji picker and returns the selected emoji
func ShowPicker() string {
	emoji, err := RunPicker()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}
	return emoji
}

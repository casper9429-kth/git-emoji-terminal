package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/casper/gitemoji"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

// Model for our example application
type model struct {
	textarea    textarea.Model
	emojiChan   chan string
	emojiPicker *tea.Program
	quitting    bool
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Type '::' to trigger emoji picker..."
	ta.Focus()

	return model{
		textarea: ta,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.textarea.SetWidth(msg.Width - 2)
		m.textarea.SetHeight(msg.Height - 4)
	}

	// Handle special case for emoji picker trigger
	if m.textarea.Value() != "" && m.emojiChan == nil && len(m.textarea.Value()) >= 2 {
		lastTwo := m.textarea.Value()[len(m.textarea.Value())-2:]
		if lastTwo == "::" {
			// Start the emoji picker
			m.emojiChan = gitemoji.EmojiPicker()

			// Start a goroutine to wait for emoji selection
			go func() {
				emoji := <-m.emojiChan
				if emoji != "" {
					// Remove the "::" trigger and add emoji
					text := m.textarea.Value()
					newText := text[:len(text)-2] + emoji + " "

					// Update textarea
					m.textarea.SetValue(newText)

					// Notify that we've processed the emoji
					m.emojiChan = nil

					// Force a refresh of the UI
					program.Send(tea.KeyMsg{Type: tea.KeyRunes})
				} else {
					// If no emoji was selected, just close the picker
					m.emojiChan = nil
				}
			}()
		}
	}

	// Handle textarea updates
	var textareaCmd tea.Cmd
	m.textarea, textareaCmd = m.textarea.Update(msg)
	cmds = append(cmds, textareaCmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	return fmt.Sprintf(
		"%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	)
}

// Clear the terminal screen
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var program *tea.Program

func main() {
	clearScreen()

	m := initialModel()
	program = tea.NewProgram(m, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}

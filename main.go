package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput" // NEW: Import textinput
	tea "github.com/charmbracelet/bubbletea"
)

// NEW: Define modes to track what the user is doing
type mode int

const (
	modeView mode = iota // Browsing the list
	modeEdit             // Typing in the text box
)

type model struct {
	tasks     []Task
	cursor    int
	mode      mode            // NEW: Track current mode
	textInput textinput.Model // NEW: The text input component
	err       error
}

func initialModel(tasks []Task) model {
	ti := textinput.New()
	ti.Placeholder = "Task name..."
	ti.CharLimit = 156
	ti.Width = 30

	return model{
		tasks:     tasks,
		cursor:    0,
		mode:      modeView, // Start in View mode
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink // NEW: Make the cursor blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// NEW: Handle updates differently based on the mode
	switch m.mode {
	case modeView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}
			case "enter", " ":
				if len(m.tasks) > 0 {
					m.tasks[m.cursor].Done = !m.tasks[m.cursor].Done
					_ = SaveTasks(m.tasks)
				}
			case "delete", "backspace":
				if len(m.tasks) > 0 {
					m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
					if m.cursor >= len(m.tasks) && m.cursor > 0 {
						m.cursor--
					}
					_ = SaveTasks(m.tasks)
				}
			// NEW: Press 'e' to enter Edit Mode
			case "e":
				if len(m.tasks) > 0 {
					m.mode = modeEdit
					m.textInput.SetValue(m.tasks[m.cursor].Title) // Pre-fill with current text
					m.textInput.Focus()
					return m, textinput.Blink
				}
			case "n":
				newTask := Task{ID: len(m.tasks) + 1, Title: "New Task (Edit me)", Done: false}
				m.tasks = append(m.tasks, newTask)
				_ = SaveTasks(m.tasks)
			}
		}

	case modeEdit:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				// Save changes
				m.tasks[m.cursor].Title = m.textInput.Value()
				m.mode = modeView
				m.textInput.Blur()
				_ = SaveTasks(m.tasks) // Write to disk
				return m, nil
			case "esc":
				// Cancel changes
				m.mode = modeView
				m.textInput.Blur()
				return m, nil
			}
		}
		// Pass keystrokes to the text input component
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, cmd
}

func (m model) View() string {
	s := "Todo List\n\n"

	for i, task := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if task.Done {
			checked = "x"
		}

		// NEW: If we are editing THIS specific task, render the input box instead of the text
		if m.mode == modeEdit && m.cursor == i {
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, m.textInput.View())
		} else {
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Title)
		}
	}

	s += "\nControls:\n"
	if m.mode == modeEdit {
		s += "Enter: Save • Esc: Cancel"
	} else {
		s += "↑/↓: Move • Space: Toggle • 'n' New • 'e': Edit • 'q': Quit • BackSpace: Delete"
	}

	return s
}

func main() {
	tasks, err := LoadTasks()
	if err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	// Use our new initialModel function
	m := initialModel(tasks)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

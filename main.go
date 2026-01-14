package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput" // NEW: Import textinput
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NEW: Define modes to track what the user is doing
type mode int

const (
	modeView mode = iota // Browsing the list
	modeEdit             // Typing in the text box
)

var (
	windowStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	doneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Strikethrough(true)

	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1).
			MarginBottom(1).
			Bold(true)
)

type model struct {
	width     int
	height    int
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
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.mode {
	case modeView:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width = msg.Width
			m.height = msg.Height
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
			case "e":
				if len(m.tasks) > 0 {
					m.mode = modeEdit
					m.textInput.SetValue(m.tasks[m.cursor].Title) // Pre-fill with current text
					m.textInput.Focus()
					return m, textinput.Blink
				}
			case "n":
				// 1. Create a blank task
				newTask := Task{ID: len(m.tasks) + 1, Title: "New Task", Done: false}
				m.tasks = append(m.tasks, newTask)

				// 2. Move cursor to the new task
				m.cursor = len(m.tasks) - 1

				// 3. Switch to edit mode immediately
				m.mode = modeEdit
				m.textInput.SetValue("")
				m.textInput.Focus()
				return m, textinput.Blink
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
	// 1. Header
	s := titleStyle.Render("TODO MANAGER") + "\n"

	// 2. Build the list
	var listContent string
	for i, task := range m.tasks {
		cursor := " "
		var taskTitle string

		// If we are editing THIS specific line, show the Input Box
		if m.mode == modeEdit && m.cursor == i {
			cursor = ">"
			// This renders the actual blinking text box
			taskTitle = m.textInput.View()
		} else {
			taskTitle = task.Title
			if task.Done {
				taskTitle = doneStyle.Render(taskTitle)
			}
			if m.cursor == i {
				cursor = ">"
				taskTitle = selectedStyle.Render(taskTitle)
			}
		}

		checked := "[ ]"
		if task.Done {
			checked = "[x]"
		}

		listContent += fmt.Sprintf("%s %s %s\n", cursor, checked, taskTitle)
	}

	// 3. Footer/Help
	helpStr := "\nn: new • e: edit • space: toggle • backspace: delete • q: quit"
	if m.mode == modeEdit {
		helpStr = "\nenter: save • esc: cancel"
	}

	help := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(helpStr)

	content := s + listContent + help

	styledContent := windowStyle.Render(content)

	return lipgloss.Place(
		m.width,         // Total width to fill
		m.height,        // Total height to fill
		lipgloss.Center, // Horizontal Alignment
		lipgloss.Center, // Vertical Alignment
		styledContent,   // Your content
	)
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

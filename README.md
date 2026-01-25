# ✅ ToDo TUI App

A sleek, keyboard-centric To-Do list application built for the terminal. Manage your tasks with style using a distraction-free interface powered by Go, Bubble Tea, and Lip Gloss.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)

## 🌟 Features

- 🎨 **Beautiful TUI** - Elegant, terminal-adaptive styling with Lip Gloss
- ⌨️ **Keyboard-Centric** - Fully navigable with intuitive keybindings
- 💾 **Persistent Storage** - Tasks are saved automatically
- 🎯 **Distraction-Free** - Minimal interface for maximum productivity
- 🔄 **Robust State Management** - Powered by Bubble Tea framework
- 🚀 **Fast & Lightweight** - Written in Go for performance
- 📱 **Cross-Platform** - Works on Linux, macOS, and Windows

## 🎬 Demo
https://github.com/user-attachments/assets/725c6194-4a64-4609-9470-bddedcc6b62e

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Terminal with 256-color support (recommended)

### Installation

#### Option 1: Using `go install`

```bash
go install github.com/Ilya-sss/ToDo-Tui-App@latest
```

#### Option 2: Clone and Build

```bash
# Clone the repository
git clone https://github.com/Ilya-sss/ToDo-Tui-App.git
cd ToDo-Tui-App

# Build the application
go build -o todo

# Run it
./todo
```

#### Option 3: Using Make

```bash
# Clone the repository
git clone https://github.com/Ilya-sss/ToDo-Tui-App.git
cd ToDo-Tui-App

# Build using make
make build

# Run the application
make run
```

## 🎮 Usage

### Basic Commands

Launch the application:
```bash
./todo
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `a` / `n` | Add a new task |
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Space` / `Enter` | Toggle task completion |
| `d` / `x` | Delete selected task |
| `e` | Edit selected task |
| `q` / `Ctrl+C` | Quit application |
| `?` | Show help menu |

### Task Management

1. **Add a task**: Press `a`, type your task, and hit `Enter`
2. **Complete a task**: Navigate to it and press `Space`
3. **Delete a task**: Navigate to it and press `d`
4. **Edit a task**: Navigate to it and press `e`

## 📁 Project Structure

```
ToDo-Tui-App/
│
├── main.go              # Application entry point
├── todo.go              # Core To-Do logic and models
├── data/                # Data storage directory
│   └── todos.json       # Persistent task storage
│
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── Makefile             # Build automation
└── README.md            # This file
```

## 🛠️ Built With

- **[Go](https://golang.org/)** - Programming language
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - TUI framework for robust state management
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for elegant terminal rendering
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - Common TUI components

## 🔧 Development

### Setting Up Development Environment

```bash
# Clone the repository
git clone https://github.com/Ilya-sss/ToDo-Tui-App.git
cd ToDo-Tui-App

# Install dependencies
go mod download

# Run in development mode
go run .
```

### Available Make Commands

```bash
make build      # Build the application
make run        # Run the application
make clean      # Clean build artifacts
make test       # Run tests
make install    # Install to $GOPATH/bin
```

### Code Structure

The application follows the Elm Architecture pattern (via Bubble Tea):

```go
type Model struct {
    todos    []Todo
    cursor   int
    selected map[int]struct{}
}

func (m Model) Init() tea.Cmd
func (m Model) Update(tea.Msg) (tea.Model, tea.Cmd)
func (m Model) View() string
```

## 📝 Configuration

Tasks are stored in `data/todos.json` by default. You can modify the data directory in the source code if needed.

### Data Format

```json
{
    {
      "id": 1,
      "text": "Your task here",
      "done": false,
    }
}
```

## 🎨 Customization

You can customize the appearance by modifying the Lip Gloss styles in the source code:

```go
// Example: Change the accent color
var accentColor = lipgloss.Color("#00ADD8")

// Example: Modify box borders
var boxStyle = lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(accentColor)
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. 🍴 Fork the repository
2. 🔀 Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. ✍️ Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. 📤 Push to the branch (`git push origin feature/AmazingFeature`)
5. 🎉 Open a Pull Request

### Ideas for Contributions

- 🏷️ Add task categories/tags
- 📅 Implement due dates
- ⭐ Add priority levels
- 🔍 Add search functionality
- 📊 Add task statistics
- 🎨 More color themes
- 🌙 Dark/light mode toggle

## 🐛 Bug Reports

Found a bug? Please open an issue with:
- Your OS and terminal
- Go version (`go version`)
- Steps to reproduce
- Expected vs actual behavior

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) - For the amazing TUI libraries
- The Go community - For excellent tooling and support
- All contributors who help improve this project

## 🌟 Show Your Support

If you find this project useful, please give it a ⭐️ on GitHub!

## 📬 Contact

**Created by [Ilya-sss](https://github.com/Ilya-sss)**

Have questions or suggestions? Feel free to [open an issue](https://github.com/Ilya-sss/ToDo-Tui-App/issues)!

---

**Stay productive from your terminal! 🚀**

package main

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	screenWidth  int
	screenHeight int
	screenStart  int

	content []string
}

func (m *model) Init() tea.Cmd {
	return nil
}

// this one works
func (m *model) View2() string {
	out := m.content[m.screenStart : m.screenStart+m.screenHeight]
	return strings.Join(out, "\n")
}

// this one does not work
func (m *model) View() string {
	var sb strings.Builder

	// uncommenting this line fixes the bug
	// sb.WriteRune('\n')

	for i, line := range m.content {
		if i >= m.screenStart && i < m.screenStart+m.screenHeight {
			sb.WriteString(line)
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.screenStart -= 1
			if m.screenStart < 0 {
				m.screenStart = 0
			}
		case "down":
			m.screenStart += 1
			if m.screenStart > len(m.content)-m.screenHeight {
				m.screenStart = len(m.content) - m.screenHeight
			}
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width
	}
	return m, nil
}

func main() {
	bytes, err := os.ReadFile("lorem_ipsum.txt")
	if err != nil {
		panic(err.Error())
	}
	lorem := string(bytes)
	model := model{
		content: strings.Split(lorem, "\n"),
	}
	p := tea.NewProgram(
		&model,
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		panic(err.Error())
	}
}

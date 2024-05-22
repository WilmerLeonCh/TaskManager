package UIMessageText

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/TaskManager/utils"
)

var messageStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#aaacb0")).
	Italic(true)

func Create(message string) {
	bubble := tea.NewProgram(initialModel(message))
	utils.Must(bubble.Run())
}

type model struct {
	message string
}

func initialModel(message string) model {
	return model{message: message}
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteRune('\n')
	b.WriteString(messageStyle.Render(m.message))
	b.WriteRune('\n')
	return b.String()
}

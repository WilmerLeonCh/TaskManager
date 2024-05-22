package UIUpdateForm

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	tasks "github.com/TaskManager/internal"
	"github.com/TaskManager/utils"
)

var (
	formTask         = tasks.MTask{}
	violet           = lipgloss.Color("57")
	darkGray         = lipgloss.Color("#767676")
	promptStyle      = lipgloss.NewStyle().Background(violet).Foreground(lipgloss.Color("229"))
	placeholderStyle = lipgloss.NewStyle().Foreground(darkGray)
)

func Create(existedTask tasks.MTask) tasks.MTask {
	bubble := tea.NewProgram(initialModel(existedTask))
	utils.Must(bubble.Run())
	return formTask
}

type model struct {
	inputs []textinput.Model
	focus  int
	err    error
}

const (
	name = iota
	description
)

func initialModel(existedTask tasks.MTask) model {
	formTask = existedTask
	var inputs = make([]textinput.Model, 2)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = fmt.Sprintf("previous: %s", existedTask.Name)
	inputs[name].PlaceholderStyle = placeholderStyle
	inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 50
	inputs[name].Validate = nameValidator

	inputs[description] = textinput.New()
	inputs[description].Placeholder = fmt.Sprintf("previous: %s", existedTask.Description)
	inputs[description].PlaceholderStyle = placeholderStyle
	inputs[description].CharLimit = 40
	inputs[description].Width = 50
	inputs[description].Validate = descriptionValidator

	return model{
		inputs: inputs,
		focus:  0,
		err:    nil,
	}
}

func nameValidator(s string) error {
	formTask.Name = s
	return nil
}

func descriptionValidator(s string) error {
	formTask.Description = s
	return nil
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.focus == len(m.inputs)-1 && m.inputs[name].Value() != "" {
				return m, tea.Quit
			}
			m.focus = nextInput(m)
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "down":
			m.focus = nextInput(m)
		case "shift+tab", "up":
			m.focus = prevInput(m)
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focus].Focus()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func nextInput(m model) int {
	return (m.focus + 1) % len(m.inputs)
}

func prevInput(m model) int {
	m.focus--
	if m.focus < 0 {
		m.focus = len(m.inputs) - 1
	}
	return m.focus
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString(promptStyle.Width(m.inputs[name].Width).Render(" Name: "))
	b.WriteRune('\n')
	b.WriteString(m.inputs[name].View())
	b.WriteRune('\n')
	b.WriteString(promptStyle.Width(m.inputs[description].Width).Render(" Description: "))
	b.WriteRune('\n')
	b.WriteString(m.inputs[description].View())
	b.WriteRune('\n')
	return b.String()
}

package UIDetailsList

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	tasks "github.com/TaskManager/internal"
	"github.com/TaskManager/ui"
	"github.com/TaskManager/utils"
)

func Create(task tasks.MTask) {
	bubble := tea.NewProgram(initialModel(task))
	utils.Must(bubble.Run())
}

type model struct {
	list []string
}

const (
	id = iota
	name
	description
	completed
	createdAt
)

func initialModel(task tasks.MTask) model {
	return model{
		list: []string{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			task.Description,
			fmt.Sprintf("%t", task.Completed),
			fmt.Sprintf("%s", task.CreatedAt),
		},
	}
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString(ui.TitleListStyle.Render(fmt.Sprintf("Task detail of # %s", m.list[id])))
	b.WriteRune('\n')

	b.WriteString(ui.ItemListStyle.Render("- Name: "))
	b.WriteString(m.list[name])
	b.WriteRune('\n')
	b.WriteString(ui.ItemListStyle.Render("- Description: "))
	b.WriteString(m.list[description])
	b.WriteRune('\n')
	b.WriteString(ui.ItemListStyle.Render("- Completed: "))
	b.WriteString(m.list[completed])
	b.WriteRune('\n')
	b.WriteString(ui.ItemListStyle.Render("- Created at: "))
	b.WriteString(m.list[createdAt])
	b.WriteRune('\n')
	return b.String()
}

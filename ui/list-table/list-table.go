package UIListTable

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/TaskManager/ui"
	"github.com/TaskManager/utils"
)

func Create(columns []table.Column, rows []table.Row) {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)

	s := table.Styles{
		Header:   ui.HeaderTableStyle,
		Selected: ui.SelectedTableStyle,
	}
	t.SetStyles(s)

	m := model{t}
	utils.Must(tea.NewProgram(m).Run())
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return ui.TableStyle.Render(m.table.View()) + "\n"
}

package table

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/evertras/bubble-table/table"
)

type Model struct {
  m table.Model
}

func New(m table.Model) tea.Model {
  return Model{
    m: m,
  }
}

func (m Model) Init() tea.Cmd {
  return m.m.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  m.m, cmd = m.m.Update(msg)
  return m, cmd
}

func (m Model) View() string {
  return m.m.View()
}


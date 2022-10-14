package table

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/bubbles/table"
)

type SetRows struct{
  rows []table.Row
}

type Model struct{
  m table.Model
}

func New(m table.Model) Model {
  return Model{
    m: m,
  }
}

func (m Model) Init() tea.Cmd {
  return nil 
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case SetRows:
    m.m.SetRows(msg.rows)
  default:
    m.m, cmd = m.m.Update(msg)
  }

  return m, cmd
}

func (m Model) View() string {
  return m.m.View()
}

func (m *Model) SetRows(r []table.Row) {
  m.m.SetRows(r)
}


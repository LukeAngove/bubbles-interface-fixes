package list

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/bubbles/list"
)

type Model struct {
  m list.Model
}

func New(m list.Model) tea.Model {
  return Model{
    m: m,
  }
}

func (m Model) Init() tea.Cmd {
  return nil 
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  m.m, cmd = m.m.Update(msg)
  return m, cmd
}

func (m Model) View() string {
  return m.m.View()
}


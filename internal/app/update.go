package app

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

type model struct{}

// Init implements [tea.Model].
func (m *model) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements [tea.Model].
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("%v", msg)
	return m, nil
}

// View implements [tea.Model].
func (m *model) View() tea.View {
	panic("unimplemented")
}

var _ tea.Model = (*model)(nil)

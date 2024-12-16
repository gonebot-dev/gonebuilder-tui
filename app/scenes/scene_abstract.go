package scenes

import tea "github.com/charmbracelet/bubbletea"

type Scene struct{}

func (s *Scene) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (s *Scene) View() string {
	return ""
}

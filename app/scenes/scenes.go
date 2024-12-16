package scenes

import tea "github.com/charmbracelet/bubbletea"

type Scene struct{}

// Name of the scene
func (s Scene) Name() string {
	return "scene"
}

// Update the scene, return the next scene and any bubbletea command.
//
// Returns nil if the scene should not change.
func (s *Scene) Update(msg tea.Msg) (*Scene, tea.Cmd) {
	return nil, nil
}

// View of the scene.
func (s *Scene) View() string {
	return ""
}

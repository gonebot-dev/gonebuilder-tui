package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Scene interface {
	// Name of the scene
	Name() string
	// Update the scene, return the next scene and any bubbletea command.
	//
	// Returns the original scene if the scene should not change.
	Update(msg tea.Msg) (Scene, tea.Cmd)
	// View of the scene.
	View() string
}

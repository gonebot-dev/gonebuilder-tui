package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
)

type initialScene struct {
	Scene
}

func (is initialScene) Name() string {
	return "[app_initial]initial_scene"
}

func (is initialScene) Update(msg tea.Msg) (Scene, tea.Cmd) {
	// TODO: Implement
	return is, nil
}

func (is initialScene) View() string {
	// TODO: Implement
	return ""
}

var InitialScene = initialScene{}

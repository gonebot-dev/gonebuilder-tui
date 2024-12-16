package initialscene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/scenes"
)

type InitialScene struct {
	scenes.Scene
}

func (is InitialScene) Name() string {
	return "[app_initial]initial_scene"
}

func (is *InitialScene) Update(msg tea.Msg) (*scenes.Scene, tea.Cmd) {
	// TODO: Implement
	return nil, nil
}

func (is *InitialScene) View() string {
	// TODO: Implement
	return ""
}

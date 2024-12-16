package initialscene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/scene"
)

type initialScene struct {
	scene.Scene
}

func (is initialScene) Update(msg tea.Msg) (scene.Scene, string, tea.Cmd) {
	// TODO: Implement
	return is, "InitialScene", nil
}

func (is initialScene) View() string {
	// TODO: Implement
	return ""
}

var InitialScene = initialScene{}

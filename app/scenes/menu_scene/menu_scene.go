package menuscene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/scene"
)

type menuScene struct {
	scene.Scene
}

func (is menuScene) Update(msg tea.Msg) (scene.Scene, string, tea.Cmd) {
	// TODO: Implement
	return is, "MenuScene", nil
}

func (is menuScene) View() string {
	// TODO: Implement
	return ""
}

var MenuScene = menuScene{}

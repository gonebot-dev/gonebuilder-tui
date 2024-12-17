package menuscene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
)

type menuScene struct {
	router.Scene
}

func (ms menuScene) Init() tea.Cmd {
	return nil
}

func (ms menuScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: Implement
	return ms, nil
}

func (ms menuScene) View() string {
	// TODO: Implement
	return ""
}

var MenuScene = menuScene{}

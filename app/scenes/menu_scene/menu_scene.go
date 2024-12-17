package menuscene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
)

type menuScene struct {
	router.Scene
}

func (ms menuScene) Name() string {
	return "MenuScene"
}

func (ms menuScene) Init() tea.Cmd {
	return nil
}

func (ms menuScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return ms, tea.Quit
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}
	return ms, nil
}

func (ms menuScene) View() string {
	// TODO: Implement
	return ""
}

var MenuScene = menuScene{}

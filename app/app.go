package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/scenes"
)

type App struct {
	CurrentScene *scenes.Scene
}

func (app App) Init() tea.Cmd {
	return nil
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return app, app.CurrentScene.Update(msg)
}

func (app App) View() string {
	return app.CurrentScene.View()
}

func CreateApp() App {
	return App{}
}

package app

import (
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/gonebot-dev/gonebuilder-tui/app/router"
	"github.com/gonebot-dev/gonebuilder-tui/app/scene"
)

type App struct {
	CurrentScene string
}

func (app App) Init() tea.Cmd {
	return nil
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	app.CurrentScene, cmd = scene.Update(app.CurrentScene, msg)
	return app, cmd
}

func (app App) View() string {
	return scene.View(app.CurrentScene)
}

func CreateApp() App {
	return App{
		CurrentScene: "InitialScene",
	}
}

package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	initialscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/initial_scene"
	menuscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/menu_scene"
)

type App struct {
	CurrentScene string
}

func (app App) Init() tea.Cmd {
	return router.Init()
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	app.CurrentScene, cmd = router.Update(app.CurrentScene, msg)
	return app, cmd
}

func (app App) View() string {
	return router.View(app.CurrentScene)
}

func init() {
	router.RegisterScene("InitialScene", initialscene.InitialScene)
	router.RegisterScene("MenuScene", menuscene.MenuScene)
}

func CreateApp() App {
	return App{
		CurrentScene: "InitialScene",
	}
}

package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	initialscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/initial_scene"
	menuscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/menu_scene"
	newbotscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/new_bot_scene"
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
	if base.WindowWidth <= 75 || base.WindowHeight <= 24 {
		widthStyle := lipgloss.NewStyle().Foreground(base.Colors.Green)
		heightStyle := lipgloss.NewStyle().Foreground(base.Colors.Green)
		if base.WindowWidth <= 75 {
			widthStyle = lipgloss.NewStyle().Foreground(base.Colors.Red)
		}
		if base.WindowHeight <= 24 {
			heightStyle = lipgloss.NewStyle().Foreground(base.Colors.Red)
		}
		return lipgloss.NewStyle().
			Width(base.WindowWidth).
			Height(base.WindowHeight).
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render(
				fmt.Sprintf(
					"Terminal size too small:\nWidth = %s, Height = %s\n\nNeeded for the application:\nWidth = %s, Height = %s",
					widthStyle.Render(fmt.Sprintf("%d", base.WindowWidth)),
					heightStyle.Render(fmt.Sprintf("%d", base.WindowHeight)),
					lipgloss.NewStyle().Foreground(base.Colors.Blue).Render("75"),
					lipgloss.NewStyle().Foreground(base.Colors.Blue).Render("24"),
				),
			)
	}
	return router.View(app.CurrentScene)
}

func init() {
	router.RegisterScene(initialscene.InitialScene)
	router.RegisterScene(menuscene.MenuScene)
	router.RegisterScene(newbotscene.NewBotScene)
}

func CreateApp() App {
	return App{
		CurrentScene: "InitialScene",
	}
}

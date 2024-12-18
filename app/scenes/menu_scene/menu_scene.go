package menuscene

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	t "github.com/gonebot-dev/gonebuilder-tui/app/translator"
)

type menuScene struct {
	router.Scene
	form *huh.Form
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
		case "ctrl+c", "esc":
			return ms, tea.Quit
		case "ctrl+l":
			if base.Lang == "en" {
				base.Lang = "zh"
			} else {
				base.Lang = "en"
			}
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}
	return ms, nil
}

func (ms menuScene) View() string {
	Header = Header.Width(base.WindowWidth)
	Footer = Footer.Width(base.WindowWidth)
	Content = Content.Width(base.WindowWidth).
		Height(base.WindowHeight - Header.GetHeight() - Footer.GetHeight())
	return MainFrame.Render(fmt.Sprintf(
		"%s\n%s\n%s",
		Header.Render(fmt.Sprintf("Current width: %d, height: %d", base.WindowWidth, base.WindowHeight)),
		Content.Render(""),
		Footer.Render(""),
	))
}

var MenuScene = menuScene{
	form: huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				TitleFunc(func() string { return t.Translate("What can I do for you?") }, base.Lang).
				OptionsFunc(
					func() []huh.Option[string] {
						return []huh.Option[string]{
							huh.NewOption(t.Translate("Create a new gonebot."), ""),
							huh.NewOption(t.Translate("Modify an existing gonebot."), ""),
							huh.NewOption(t.Translate("Manage .env configurations."), ""),
							huh.NewOption(t.Translate("Explore plugin repository."), ""),
							huh.NewOption(t.Translate("Explore adapter repository."), ""),
							huh.NewOption(t.Translate("Exit the application."), ""),
						}
					},
					base.Lang,
				),
		),
	),
}

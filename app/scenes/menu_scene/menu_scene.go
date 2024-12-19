package menuscene

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
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
	return ms.form.Init()
}

func (ms menuScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return ms, tea.Quit
		case tea.KeyCtrlL:
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

	var cmds []tea.Cmd
	form, cmd := ms.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		cmds = append(cmds, cmd)
		ms.form = f
	}

	if ms.form.State == huh.StateCompleted {
		base.SelectedAction = ms.form.GetInt("action")
		if os.Getenv("DEBUG") == "true" {
			return ms, tea.Quit
		}
		// TODO: Add jump scenes
	}

	return ms, tea.Batch(cmds...)
}

func (ms menuScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth - 32)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 2)
	base.FormStyle = base.FormStyle.Width(min(base.WindowWidth-8, 54))

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n%s\n%s",
		base.Header.Render("GoneBuilder"),
		base.Content.Render(
			base.FormStyle.Render(ms.form.WithHeight(10).View()),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			base.Footer.Render(
				fmt.Sprintf("%s%s%s%s",
					base.FooterTitle.Render("Exit"),
					base.FooterText.Render("Ctrl+C"),
					base.FooterTitle.Render(t.Translate("让我们说中文")),
					base.FooterText.Render("Ctrl+L"),
				),
			),
			base.FooterCopyright.Render("Copyright © 2024 gonebot-dev"),
		),
	))
}

var MenuScene = menuScene{
	form: huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Key("action").
				TitleFunc(func() string {
					return lipgloss.NewStyle().
						Bold(true).
						Render(t.Translate("What can I do for you?"))
				}, &base.Lang).
				DescriptionFunc(func() string { return t.Translate("Select an option to continue...") + "\n" }, &base.Lang).
				OptionsFunc(
					func() []huh.Option[int] {
						return []huh.Option[int]{
							huh.NewOption(t.Translate("Create a new gonebot."), base.NewBot),
							huh.NewOption(t.Translate("Modify an existing gonebot."), base.EditBot),
							huh.NewOption(t.Translate("Manage .env configurations."), base.DotEnv),
							huh.NewOption(t.Translate("Explore plugin repository."), base.Plugins),
							huh.NewOption(t.Translate("Explore adapter repository."), base.Adapters),
							huh.NewOption(t.Translate("Exit the application."), base.ExitApp),
						}
					},
					&base.Lang,
				),
		),
	),
}

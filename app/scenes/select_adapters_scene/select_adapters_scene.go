package selectadaptersscene

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	t "github.com/gonebot-dev/gonebuilder-tui/app/translator"
)

type selectAdaptersScene struct {
	router.Scene
	form *huh.Form
}

func (s selectAdaptersScene) Name() string {
	return "SelectAdaptersScene"
}

func (s selectAdaptersScene) GetEmits() map[string]string {
	return map[string]string{}
}

func (s selectAdaptersScene) Init() tea.Cmd {
	return nil
}

func (s selectAdaptersScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyCtrlD:
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
	return s, nil
}

func (s selectAdaptersScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth - 32)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 2)

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n%s\n%s",
		base.Header.Render("GoneBuilder"),
		base.Content.Render(
			"",
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			base.Footer.Render(
				fmt.Sprintf("%s%s%s%s",
					base.FooterTitle.Render("Exit"),
					base.FooterText.Render("Ctrl+C"),
					base.FooterTitle.Render(t.Translate("让我们说中文")),
					base.FooterText.Render("Ctrl+D"),
				),
			),
			base.FooterCopyright.Render("Copyright © 2024 gonebot-dev"),
		),
	))
}

var SelectAdaptersScene = selectAdaptersScene{}

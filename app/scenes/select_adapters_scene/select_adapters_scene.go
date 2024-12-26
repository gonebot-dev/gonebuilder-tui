package selectadaptersscene

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
	t "github.com/gonebot-dev/gonebuilder-tui/app/utils/translator"
)

type selectAdaptersScene struct {
	router.Scene
	adapters list.Model
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

var syncing = true

func (s selectAdaptersScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if api.Finished == true && api.CurrentCommit.SHA == "" {
		cmds = append(cmds, s.adapters.ToggleSpinner())
		go api.SyncRepo()
	}
	if syncing == true && api.Finished == true && api.CurrentCommit.SHA != "" {
		syncing = false
		s.adapters.SetItems(api.Adapters)
		cmds = append(cmds, s.adapters.ToggleSpinner())
	}
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
		case tea.KeyCtrlR:
			if api.Finished {
				syncing = true
				api.CurrentCommit.SHA = ""
				cmds = append(cmds, s.adapters.ToggleSpinner())
			}
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}
	if syncing {
		s.adapters.Title = t.Translate("Syncing Repository...")
	} else {
		s.adapters.Title = t.Translate("Select Adapters...")
	}
	var cmd tea.Cmd
	s.adapters, cmd = s.adapters.Update(msg)
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s selectAdaptersScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth - 32)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 6).AlignHorizontal(lipgloss.Left)
	base.FormStyle = base.FormStyle.Width((base.WindowWidth - 8) / 2).
		Height(base.WindowHeight - 8)
	s.adapters.SetSize((base.WindowWidth-8)/2, base.WindowHeight-8)

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n%s\n%s",
		base.Header.Render("GoneBuilder"),
		base.Content.Padding(2, 2).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				base.BasicStyle.Width((base.WindowWidth-4)/2).
					Height(base.WindowHeight-8).
					PaddingRight(2).PaddingTop(1).
					Render(s.adapters.View()),
				base.FormStyle.Render(""),
			),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			base.Footer.Render(
				fmt.Sprintf("%s%s%s%s%s%s",
					base.FooterTitle.Render(t.Translate("Exit")),
					base.FooterText.Render("Ctrl+C"),
					base.FooterTitle.Render(t.Translate("让我们说中文")),
					base.FooterText.Render("Ctrl+D"),
					base.FooterTitle.Render(t.Translate("Refresh")),
					base.FooterText.Render("Ctrl+R"),
				),
			),
			base.FooterCopyright.Render("Copyright © 2024 gonebot-dev"),
		),
	))
}

var SelectAdaptersScene = selectAdaptersScene{
	adapters: list.New([]list.Item{}, list.NewDefaultDelegate(), (base.WindowWidth-4)/2, base.WindowHeight-8),
}

func init() {
	SelectAdaptersScene.adapters.Title = t.Translate("Syncing Repository...")
	SelectAdaptersScene.adapters.SetSpinner(spinner.Line)
	SelectAdaptersScene.adapters.SetShowPagination(true)
	SelectAdaptersScene.adapters.SetShowStatusBar(true)

	SelectAdaptersScene.adapters.KeyMap.Quit = key.NewBinding(key.WithDisabled())
}

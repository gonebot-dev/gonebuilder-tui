package selectpluginsscene

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	selectedlist "github.com/gonebot-dev/gonebuilder-tui/app/components/selected_list"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
	t "github.com/gonebot-dev/gonebuilder-tui/app/utils/translator"
)

type selectPluginsScene struct {
	router.Scene
	plugins list.Model
}

func (s selectPluginsScene) Name() string {
	return "SelectPluginsScene"
}

func (s selectPluginsScene) Init() tea.Cmd {
	return nil
}

func (s selectPluginsScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if api.Finished && api.CurrentCommit.SHA == "" {
		cmds = append(cmds, s.plugins.ToggleSpinner())
		api.Finished = false
		go api.SyncRepo()
	}
	if base.RepoSyncing && api.Finished && api.CurrentCommit.SHA != "" {
		base.RepoSyncing = false
		s.plugins.SetItems(api.Plugins)
		cmds = append(cmds, s.plugins.ToggleSpinner())
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyCtrlF:
			if base.Lang == "en" {
				base.Lang = "zh"
			} else {
				base.Lang = "en"
			}
		case tea.KeyCtrlR:
			if api.Finished {
				base.RepoSyncing = true
				api.CurrentCommit.SHA = ""
				cmds = append(cmds, s.plugins.ToggleSpinner())
				go api.SyncRepo()
				selectedlist.SelectedList.SelectedAdapters = make([]list.Item, 0)
				selectedlist.SelectedList.SelectedPlugins = make([]list.Item, 0)
				selectedlist.SelectedList.AdaptersList.SetItems(selectedlist.SelectedList.SelectedAdapters)
				selectedlist.SelectedList.PluginsList.SetItems(selectedlist.SelectedList.SelectedPlugins)
			}
		case tea.KeyTab:
			switch selectedlist.SelectedList.Focus {
			case "none":
				if len(selectedlist.SelectedList.SelectedAdapters) > 0 {
					selectedlist.SelectedList.Focus = "adapters"
				} else if len(selectedlist.SelectedList.SelectedPlugins) > 0 {
					selectedlist.SelectedList.Focus = "plugins"
				}
			case "adapters":
				if len(selectedlist.SelectedList.PluginsList.Items()) > 0 {
					selectedlist.SelectedList.Focus = "plugins"
				} else {
					selectedlist.SelectedList.Focus = "none"
				}
			case "plugins":
				selectedlist.SelectedList.Focus = "none"
			}
		case tea.KeyEnter:
			if selectedlist.SelectedList.Focus == "none" && len(s.plugins.Items()) > 0 {
				cmds = append(cmds, selectedlist.SelectedList.PluginsList.InsertItem(
					len(selectedlist.SelectedList.PluginsList.Items()),
					s.plugins.SelectedItem(),
				))
				selectedlist.SelectedList.SelectedPlugins = append(
					selectedlist.SelectedList.SelectedPlugins,
					s.plugins.SelectedItem(),
				)
				s.plugins.RemoveItem(s.plugins.Index())
			} else if selectedlist.SelectedList.Focus == "plugins" && len(selectedlist.SelectedList.SelectedPlugins) > 0 {
				cmds = append(cmds, s.plugins.InsertItem(
					s.plugins.Index(),
					selectedlist.SelectedList.PluginsList.SelectedItem(),
				))
				index := selectedlist.SelectedList.PluginsList.Index()
				selectedlist.SelectedList.SelectedPlugins = append(
					selectedlist.SelectedList.SelectedPlugins[:index],
					selectedlist.SelectedList.SelectedPlugins[index+1:]...,
				)
				selectedlist.SelectedList.PluginsList.RemoveItem(index)
			}
		case tea.KeyCtrlLeft:
			cmds = append(cmds, router.NextScene("SelectAdaptersScene"))
		case tea.KeyCtrlRight:
			// TODO: Jump to the next scene
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}
	if base.RepoSyncing {
		s.plugins.Title = t.Translate("Syncing Repository...")
	} else {
		s.plugins.Title = t.Translate("Select Plugins...")
	}
	var cmd tea.Cmd
	if selectedlist.SelectedList.Focus == "none" {
		s.plugins, cmd = s.plugins.Update(msg)
		cmds = append(cmds, cmd)
	}
	model, cmd := selectedlist.SelectedList.Update(msg)
	selectedlist.SelectedList = model.(selectedlist.SelectedListModel)
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s selectPluginsScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 6).AlignHorizontal(lipgloss.Left)
	s.plugins.SetSize((base.WindowWidth-8)/3*2, base.WindowHeight-14)

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n\n\n%s\n\n\n%s",
		base.Header.Render("GoneBuilder - Copyright © 2024 gonebot-dev"),
		base.Content.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				base.FormStyle.Width((base.WindowWidth-4)/3*2).
					Height(base.WindowHeight-8).
					Render(s.plugins.View()),
				selectedlist.SelectedList.View(),
			),
		),
		base.Footer.Render(
			fmt.Sprintf("%s%s%s%s%s%s%s%s",
				base.FooterTitle.Render(t.Translate("Exit")),
				base.FooterText.Render("Ctrl+C"),
				base.FooterTitle.Render(t.Translate("Refresh")),
				base.FooterText.Render("Ctrl+R"),
				base.FooterTitle.Render(t.Translate("Next")),
				base.FooterText.Render("Ctrl+Right"),
				base.FooterTitle.Render(t.Translate("Prev")),
				base.FooterText.Render("Ctrl+Left"),
			),
		),
	))
}

var SelectPluginsScene = selectPluginsScene{
	plugins: list.New([]list.Item{}, list.NewDefaultDelegate(), (base.WindowWidth-4)/2, base.WindowHeight-8),
}

func init() {
	SelectPluginsScene.plugins.Title = t.Translate("Syncing Repository...")
	SelectPluginsScene.plugins.SetSpinner(spinner.Line)
	SelectPluginsScene.plugins.SetShowPagination(true)
	SelectPluginsScene.plugins.SetShowStatusBar(true)

	SelectPluginsScene.plugins.KeyMap.Quit = key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String()),
		key.WithHelp(tea.KeyTab.String(), "switch focus"),
	)
}

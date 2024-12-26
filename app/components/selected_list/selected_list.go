package selectedlist

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
	t "github.com/gonebot-dev/gonebuilder-tui/app/utils/translator"
)

type adaptersDelegate struct{}

func (i adaptersDelegate) Height() int                               { return 1 }
func (i adaptersDelegate) Spacing() int                              { return 0 }
func (i adaptersDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (i adaptersDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	it, ok := item.(api.AdapterInfo)
	if !ok {
		return
	}

	str := fmt.Sprintf("・ %s", it.Title())

	var fn func(s ...string) string
	if index == m.Index() && SelectedList.Focus == "adapters" {
		fn = list.NewDefaultItemStyles().SelectedTitle.Render
	} else {
		fn = list.NewDefaultItemStyles().NormalTitle.Render
	}

	fmt.Fprint(w, fn(str))
}

type pluginsDelegate struct{}

func (i pluginsDelegate) Height() int                               { return 1 }
func (i pluginsDelegate) Spacing() int                              { return 0 }
func (i pluginsDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (i pluginsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	it, ok := item.(api.PluginInfo)
	if !ok {
		return
	}

	str := fmt.Sprintf("・ %s", it.Title())

	var fn func(s ...string) string
	if index == m.Index() && SelectedList.Focus == "plugins" {
		fn = list.NewDefaultItemStyles().SelectedTitle.Render
	} else {
		fn = list.NewDefaultItemStyles().NormalTitle.Render
	}

	fmt.Fprint(w, fn(str))
}

type SelectedListModel struct {
	tea.Model
	SelectedAdapters []list.Item
	SelectedPlugins  []list.Item
	AdaptersList     list.Model
	PluginsList      list.Model
	Focus            string
}

func (s SelectedListModel) Init() tea.Cmd {
	return nil
}

func (s SelectedListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	s.AdaptersList.Title = t.Translate("Selected Adapters")
	s.PluginsList.Title = t.Translate("Selected Plugins")

	if s.Focus == "adapters" {
		s.AdaptersList, cmd = s.AdaptersList.Update(msg)
	} else if s.Focus == "plugins" {
		s.PluginsList, cmd = s.PluginsList.Update(msg)
	}

	return s, cmd
}

func (s SelectedListModel) View() string {
	s.AdaptersList.SetSize(base.WindowWidth-8-(base.WindowWidth-8)/3*2, base.WindowHeight-14-(base.WindowHeight-14)/3*2)
	s.PluginsList.SetSize(base.WindowWidth-8-(base.WindowWidth-8)/3*2, (base.WindowHeight-14)/3*2)
	return base.FormStyle.Width(base.WindowWidth - 8 - (base.WindowWidth-8)/3*2).
		Height(base.WindowHeight - 8).
		Render(
			fmt.Sprintf(
				"%s\n\n%s",
				s.AdaptersList.View(),
				s.PluginsList.View(),
			),
		)
}

var SelectedList = SelectedListModel{
	SelectedAdapters: make([]list.Item, 0),
	SelectedPlugins:  make([]list.Item, 0),
	AdaptersList:     list.New([]list.Item{}, adaptersDelegate{}, 0, 0),
	PluginsList:      list.New([]list.Item{}, pluginsDelegate{}, 0, 0),
	Focus:            "none",
}

func init() {
	SelectedList.AdaptersList.Title = t.Translate("Selected Adapters")
	SelectedList.PluginsList.Title = t.Translate("Selected Plugins")

	SelectedList.AdaptersList.SetShowTitle(true)
	SelectedList.AdaptersList.SetShowPagination(true)
	SelectedList.AdaptersList.SetShowStatusBar(true)
	SelectedList.AdaptersList.SetShowHelp(false)

	SelectedList.PluginsList.SetShowTitle(true)
	SelectedList.PluginsList.SetShowPagination(true)
	SelectedList.PluginsList.SetShowStatusBar(true)
	SelectedList.PluginsList.SetShowHelp(false)

	SelectedList.AdaptersList.KeyMap.Quit = key.NewBinding(key.WithDisabled())
	SelectedList.PluginsList.KeyMap.Quit = key.NewBinding(key.WithDisabled())
}

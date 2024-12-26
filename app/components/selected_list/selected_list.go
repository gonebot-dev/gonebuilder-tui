package selectedlist

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type selectedList struct {
	tea.Model
	SelectedAdapters []list.Item
	SelectedPlugins  []list.Item
	adapterslist     list.Model
	pluginslist      list.Model
}

var SelectedList = selectedList{
	SelectedAdapters: make([]list.Item, 0),
	SelectedPlugins:  make([]list.Item, 0),
}

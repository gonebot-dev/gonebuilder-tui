package newbotscene

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	t "github.com/gonebot-dev/gonebuilder-tui/app/utils/translator"
)

type newBotScene struct {
	router.Scene
	form       *huh.Form
	filepicker *huh.FilePicker
	emits      map[string]string
}

func (s newBotScene) Name() string {
	return "NewBotScene"
}

func (s newBotScene) GetEmits() map[string]string {
	return s.emits
}

func (s newBotScene) Init() tea.Cmd {
	return tea.Batch(s.form.Init())
}

func (s newBotScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}

	s.filepicker.Title(t.Translate("Select a folder...")).
		Description(t.Translate("We will create your bot folder here."))

	var cmds []tea.Cmd
	form, cmd := s.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		cmds = append(cmds, cmd)
		s.form = f
	}

	if s.form.State == huh.StateCompleted {
		s.emits["bot_name"] = s.form.GetString("name")
		s.emits["bot_version"] = s.form.GetString("version")
		s.emits["bot_description"] = s.form.GetString("description")
		s.emits["bot_folder"] = s.form.GetString("folder")

		return router.GetScene("SelectAdaptersScene")
	}

	return s, tea.Batch(cmds...)
}

func (s newBotScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 2)
	base.FormStyle = base.FormStyle.Width(min(base.WindowWidth-8, 60))

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n%s\n%s",
		base.Header.Render("GoneBuilder - Copyright © 2024 gonebot-dev"),
		base.Content.Render(
			base.FormStyle.Render(s.form.WithHeight(16).View()),
		),
		base.Footer.Render(
			fmt.Sprintf("%s%s%s%s",
				base.FooterTitle.Render(t.Translate("Exit")),
				base.FooterText.Render("Ctrl+C"),
				base.FooterTitle.Render(t.Translate("让我们说中文")),
				base.FooterText.Render("Ctrl+F"),
			),
		),
	))
}

var NewBotScene = newBotScene{
	emits: make(map[string]string),
}

func init() {
	currentDir, _ := os.Getwd()
	NewBotScene.filepicker = huh.NewFilePicker().
		Key("folder").
		Title(t.Translate("Select a folder...")).
		Description(t.Translate("We will create your bot folder here.")).
		DirAllowed(true).
		FileAllowed(false).
		ShowPermissions(false).
		ShowHidden(false).
		ShowSize(false).
		CurrentDirectory(currentDir).
		Picking(true)
	NewBotScene.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				TitleFunc(func() string { return t.Translate("Enter bot name:") }, &base.Lang).
				PlaceholderFunc(func() string { return t.Translate("Name of your bot.") }, &base.Lang),
			huh.NewInput().
				Key("version").
				TitleFunc(func() string { return t.Translate("Enter bot version:") }, &base.Lang).
				PlaceholderFunc(func() string { return t.Translate("Version of your bot.") }, &base.Lang),
			huh.NewInput().
				Key("description").
				TitleFunc(func() string { return t.Translate("Enter bot description:") }, &base.Lang).
				PlaceholderFunc(func() string { return t.Translate("A short description of your bot.") }, &base.Lang),
			NewBotScene.filepicker,
		),
	)

	// Fix filepicker keymap
	keyMap := huh.NewDefaultKeyMap()
	keyMap.FilePicker = huh.FilePickerKeyMap{
		GoToTop:  key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "first"), key.WithDisabled()),
		GoToLast: key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "last"), key.WithDisabled()),
		PageUp:   key.NewBinding(key.WithKeys("K", "pgup"), key.WithHelp("pgup", "page up"), key.WithDisabled()),
		PageDown: key.NewBinding(key.WithKeys("J", "pgdown"), key.WithHelp("pgdown", "page down"), key.WithDisabled()),
		Back:     key.NewBinding(key.WithKeys("h", "backspace", "left", "esc"), key.WithHelp("h", "back"), key.WithDisabled()),
		Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"), key.WithDisabled()),
		Up:       key.NewBinding(key.WithKeys("up", "k", "ctrl+k", "ctrl+p"), key.WithHelp("↑", "up"), key.WithDisabled()),
		Down:     key.NewBinding(key.WithKeys("down", "j", "ctrl+j", "ctrl+n"), key.WithHelp("↓", "down"), key.WithDisabled()),

		Open:   key.NewBinding(key.WithKeys("l", "right", "enter"), key.WithHelp("→", "open")),
		Close:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("←", "back"), key.WithDisabled()),
		Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
		Next:   key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next")),
		Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	}
	NewBotScene.form.WithKeyMap(keyMap)
}

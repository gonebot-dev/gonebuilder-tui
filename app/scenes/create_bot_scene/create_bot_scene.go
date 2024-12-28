package createbotscene

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	selectedlist "github.com/gonebot-dev/gonebuilder-tui/app/components/selected_list"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
	botcreator "github.com/gonebot-dev/gonebuilder-tui/app/utils/bot_creator"
	t "github.com/gonebot-dev/gonebuilder-tui/app/utils/translator"
)

type createBotScene struct {
	router.Scene
	confirm     *huh.Confirm
	confirmForm *huh.Form
}

func (s createBotScene) Name() string {
	return "CreateBotScene"
}

func (s createBotScene) Init() tea.Cmd {
	return s.confirmForm.Init()
}

func (s createBotScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyCtrlF:
			base.Lang = base.IfElse(base.Lang == "en", "zh", "en")
		}
	case tea.WindowSizeMsg:
		base.WindowHeight = msg.Height
		base.WindowWidth = msg.Width
	}

	s.confirm.Affirmative(t.Translate("Affirmative!")).Negative(t.Translate("Negative."))

	var cmd tea.Cmd
	form, cmd := s.confirmForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		cmds = append(cmds, cmd)
		s.confirmForm = f
	}

	model, cmd := selectedlist.SelectedList.Update(msg)
	selectedlist.SelectedList = model.(selectedlist.SelectedListModel)
	cmds = append(cmds, cmd)

	if s.confirmForm.State == huh.StateCompleted {
		if s.confirmForm.GetBool("confirm") {
			base.PostFunc = func() (err error) {
				err = botcreator.CreateBot(
					base.BotFolder, base.BotName,
					base.BotVersion, base.BotDesc,
					&selectedlist.SelectedList.SelectedAdapters,
					&selectedlist.SelectedList.SelectedPlugins,
				)
				if err == nil {
					fmt.Printf("\nBot created successfully!\n")
					fmt.Printf("To run your bot, use the following command:\n")
					fmt.Printf("\n\tcd %s\n", strings.ReplaceAll(filepath.Join(base.BotFolder, base.BotName), "\\", "/"))
					fmt.Printf("\tgo run %s\n", botcreator.FormatName(base.BotName))
				}
				return
			}
		}
		return s, tea.Quit
	}

	return s, tea.Batch(cmds...)
}

func (s createBotScene) View() string {
	base.Header = base.Header.Width(base.WindowWidth)
	base.Footer = base.Footer.Width(base.WindowWidth)
	base.Content = base.Content.Width(base.WindowWidth).
		Height(base.WindowHeight - 6).AlignHorizontal(lipgloss.Left)
	s.confirmForm.WithWidth((base.WindowWidth - 8) / 3 * 2).WithHeight(5)

	return base.MainFrame.Render(fmt.Sprintf(
		"%s\n\n\n%s\n\n\n%s",
		base.Header.Render("GoneBuilder - Copyright © 2024 gonebot-dev"),
		base.Content.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				base.FormStyle.Width((base.WindowWidth-4)/3*2).
					Height(base.WindowHeight-8).
					AlignVertical(lipgloss.Center).
					Render(s.confirmForm.View()),
				selectedlist.SelectedList.View(),
			),
		),
		base.Footer.Render(
			fmt.Sprintf("%s%s%s%s%s%s",
				base.FooterTitle.Render(t.Translate("Exit")),
				base.FooterText.Render("Ctrl+C"),
				base.FooterTitle.Render(t.Translate("让我们说中文")),
				base.FooterText.Render("Ctrl+F"),
				base.FooterTitle.Render(t.Translate("Refresh")),
				base.FooterText.Render("Ctrl+R"),
			),
		),
	))
}

var CreateBotScene = createBotScene{
	confirm: huh.NewConfirm().
		TitleFunc(func() string {
			return t.Translate("Are you sure to create bot with these components?")
		}, &base.Lang).
		DescriptionFunc(func() string {
			return t.Translate("This will create or replace a folder in selected folder.")
		}, &base.Lang).
		Affirmative(t.Translate("Affirmative!")).
		Negative(t.Translate("No.")).
		Key("confirm"),
}

func init() {
	CreateBotScene.confirmForm = huh.NewForm(huh.NewGroup(CreateBotScene.confirm))
}

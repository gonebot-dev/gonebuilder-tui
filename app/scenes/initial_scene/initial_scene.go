package initialscene

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/gonebot-dev/gonebuilder-tui/app/router"
)

type initialScene struct {
	// Implement
	router.Scene

	// Options
	options struct {
		tickInterval    time.Duration
		loadingDuration time.Duration
	}

	// Cmponents
	progress progress.Model

	// Variables
	loadingElapsed time.Duration
	loadingPercent float64
}

func (is initialScene) Name() string {
	return "InitialScene"
}

func (is initialScene) GetEmits() map[string]string {
	return map[string]string{}
}

type tickMsg struct{}

func (is initialScene) Tick(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (is initialScene) Init() tea.Cmd {
	return tea.Batch(is.progress.Init(), is.Tick(is.options.tickInterval))
}

func (is initialScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return is, tea.Quit
		}
	case tea.WindowSizeMsg:
		base.WindowWidth = msg.Width
		base.WindowHeight = msg.Height
	case tickMsg:
		if is.loadingElapsed >= is.options.loadingDuration {
			if os.Getenv("DEBUG") == "true" {
				return is, tea.Quit
			} else {
				return router.GetScene("MenuScene")
			}
		}
		is.loadingElapsed += is.options.tickInterval
		is.loadingPercent = ease.OutBounce(
			float64(is.loadingElapsed) / float64(is.options.loadingDuration),
		)
		return is, is.Tick(is.options.tickInterval)
	}
	return is, nil
}

func (is initialScene) View() string {
	is.progress.Width = base.WindowWidth - MainFrame.Padding
	return MainFrame.Style.
		Width(base.WindowWidth).
		Height(base.WindowHeight).
		Render(
			fmt.Sprintf(
				"%s\n\n\n%s",
				MainFrame.BannerStyle.Render(MainFrame.Banner),
				is.progress.ViewAs(is.loadingPercent),
			),
		)
}

var InitialScene = initialScene{
	progress: progress.New(
		progress.WithDefaultGradient(),
		progress.WithoutPercentage(),
	),
	options: struct {
		tickInterval    time.Duration
		loadingDuration time.Duration
	}{
		tickInterval:    time.Millisecond * 16,
		loadingDuration: time.Second * 2,
	},
}

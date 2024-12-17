package initialscene

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
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
	sceneWidth     int
	sceneHeight    int
	loadingElapsed time.Duration
	loadingPercent float64
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
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return is, tea.Quit
		}
	case tea.WindowSizeMsg:
		is.sceneWidth = msg.Width
		is.sceneHeight = msg.Height
		is.progress.Width = msg.Width - MainFrame.Padding
	case tickMsg:
		if is.loadingElapsed >= is.options.loadingDuration {
			return is, tea.Quit
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
	return MainFrame.Style.
		Width(is.sceneWidth).
		Height(is.sceneHeight).
		Render(
			fmt.Sprintf("%s\n\n\n%s", Banner, is.progress.ViewAs(is.loadingPercent)),
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

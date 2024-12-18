package menuscene

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

var basic = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	Foreground(base.Colors.Primary).
	Bold(true)

var (
	Header = basic.Height(1).
		Background(base.Colors.Gray)
	Footer = basic.Height(1).
		Background(base.Colors.Gray)
	Content = basic.
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
	MainFrame = basic.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Foreground(base.Colors.Primary).
			Bold(true)
)

package menuscene

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

var basic = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	Foreground(base.Colors.PrimaryReverse)
var basicText = basic.
	Foreground(base.Colors.PrimaryReverse).
	Bold(true).
	Padding(0, 2)

var (
	Header = basic.Height(1).
		Background(base.Colors.Gray).
		MarginTop(2).
		AlignHorizontal(lipgloss.Center).
		Bold(true)
	Footer = basic.Height(1).
		Background(base.Colors.Gray).
		MarginBottom(2).
		Bold(true)
	Content = basic.
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
	FormStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(base.Colors.Lavender).
			Padding(2, 4)
	MainFrame = basic.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Bold(true)
	FooterTitle     = basicText.Background(base.Colors.Lavender)
	FooterText      = basicText.Background(base.Colors.Gray)
	FooterCopyright = basicText.Background(base.Colors.Yellow).MarginBottom(2)
)

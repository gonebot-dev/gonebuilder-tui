package base

import "github.com/charmbracelet/lipgloss"

var basic = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	Foreground(Colors.PrimaryReverse)
var basicText = basic.
	Foreground(Colors.PrimaryReverse).
	Bold(true).
	Padding(0, 2)

var (
	Header = basic.Height(1).
		Background(Colors.Gray).
		AlignHorizontal(lipgloss.Center).
		Bold(true)
	Footer = basic.Height(1).
		Background(Colors.Gray).
		Bold(true)
	Content = basic.
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
	FormStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Colors.Lavender).
			Padding(2, 4)
	MainFrame = basic.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Bold(true)
	FooterTitle     = basicText.Background(Colors.Lavender)
	FooterText      = basicText.Background(Colors.Gray)
	FooterCopyright = basicText.Background(Colors.Yellow)
)

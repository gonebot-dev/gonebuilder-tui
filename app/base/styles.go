package base

import "github.com/charmbracelet/lipgloss"

var BasicStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	Foreground(Colors.PrimaryReverse)
var BasicTextStyle = BasicStyle.
	Foreground(Colors.PrimaryReverse).
	Bold(true).
	Padding(0, 2)

var (
	Header = BasicStyle.Height(1).
		Background(Colors.Yellow).
		AlignHorizontal(lipgloss.Center).
		Bold(true)
	Footer = BasicStyle.Height(1).
		Background(Colors.Gray).
		Bold(true)
	Content = BasicStyle.
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
	FormStyle = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Left).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Colors.Lavender).
			Padding(2, 4)
	MainFrame = BasicStyle.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Bold(true)
	FooterTitle = BasicTextStyle.Background(Colors.Lavender)
	FooterText  = BasicTextStyle.Background(Colors.Gray)
)

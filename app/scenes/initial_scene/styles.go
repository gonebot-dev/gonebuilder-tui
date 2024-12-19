package initialscene

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

var MainFrame = struct {
	Style       lipgloss.Style
	Padding     int
	Banner      string
	BannerStyle lipgloss.Style
}{
	Style: lipgloss.NewStyle().
		Bold(true).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center),
	Padding: 8,
	Banner: "\n" +
		" ██████╗  ██████╗ ███╗   ██╗███████╗██████╗  ██████╗ ████████╗\n" +
		"██╔════╝ ██╔═══██╗████╗  ██║██╔════╝██╔══██╗██╔═══██╗╚══██╔══╝\n" +
		"██║  ███╗██║   ██║██╔██╗ ██║█████╗  ██████╔╝██║   ██║   ██║   \n" +
		"██║   ██║██║   ██║██║╚██╗██║██╔══╝  ██╔══██╗██║   ██║   ██║   \n" +
		"╚██████╔╝╚██████╔╝██║ ╚████║███████╗██████╔╝╚██████╔╝   ██║   \n" +
		" ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚══════╝╚═════╝  ╚═════╝    ╚═╝   \n" +
		"\n",
	BannerStyle: lipgloss.NewStyle().Foreground(base.Colors.GoBlue),
}

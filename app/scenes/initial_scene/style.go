package initialscene

import "github.com/charmbracelet/lipgloss"

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
	BannerStyle: lipgloss.NewStyle().Foreground(
		lipgloss.CompleteAdaptiveColor{
			Light: lipgloss.CompleteColor{
				ANSI:      "14",
				ANSI256:   "45",
				TrueColor: "#367b99",
			},
			Dark: lipgloss.CompleteColor{
				ANSI:      "6",
				ANSI256:   "31",
				TrueColor: "#367b99",
			},
		},
	),
}
